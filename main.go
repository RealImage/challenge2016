package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"chng2016/pkg/datasource"
	localcache "chng2016/pkg/datasource/cache"
	localdb "chng2016/pkg/datasource/localDB"
	trie "chng2016/pkg/datasource/trie"
	"chng2016/pkg/handlers"
	"chng2016/pkg/routes"
	"chng2016/pkg/utils"
	"chng2016/pkg/validation"
	"chng2016/resouces"

	"github.com/gin-gonic/gin"
)

var (
	errRest        = make(chan error)
	done           = make(chan struct{})
	restServerPort string
	helpFlag       bool
)

func init() {
	flag.BoolVar(&helpFlag, "help", false, "show usage and exit")
	flag.StringVar(&restServerPort, "port", ":8001", "rest server port")
}

func parseFlags() {
	flag.Parse()
	if helpFlag {
		flag.Usage()
		os.Exit(0)
	}
}

func handleInterrupts() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	sig := <-interrupt
	log.Println("sig : ", sig)
	done <- struct{}{}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	parseFlags()
	go handleInterrupts()

	// local db and cache declaration for loading csv file
	trie := trie.NewTrie()
	localDB := localdb.NewLocalDBClient()
	localCache := localcache.NewCacheClient()

	clientDataStore := datasource.NewDatasourceClient(localCache, localDB, trie)
	l := resouces.NewLoader(clientDataStore)
	err := l.LoadCSV()
	if err != nil {
		log.Fatal(err)
	}
	restServer := gin.Default()
	u := utils.NewAppUtil(localDB)
	v := validation.NewValidation(u)
	v.RegistorCustomValidationFunction()
	h := handlers.NewDistributorHandler(clientDataStore, v, u)
	r := routes.NewRoutes(h)
	routes.AttachRoutes(restServer, r)

	go func() {
		errRest <- restServer.Run(":8001")
	}()

	select {
	case err := <-errRest:
		log.Println("error running s : ", err)
	case <-done:
		log.Println("down server")
	}

	time.AfterFunc(1*time.Second, func() {
		close(errRest)
		close(done)
	})
}
