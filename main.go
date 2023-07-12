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

// // Online Go compiler to run Golang program online
// // Print "Hello World!" message
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strings"
// )

// // Structure to represent a node in the Trie
// type TrieNode struct {
// 	children map[string]*TrieNode
// }

// // Structure to represent the distributor's permissions
// type DistributorPermissions struct {
// 	Include []string `json:"include"`
// 	Exclude []string `json:"exclude"`
// }

// // Structure to store the distributor and its permissions
// type Distributor struct {
// 	ID          string
// 	Permissions *DistributorPermissions
// 	TrieRoot    *TrieNode
// }

// // Global map to store distributors and their permissions
// var distributorMap = make(map[string]*Distributor)

// // Function to add permissions for a distributor
// func addPermission(w http.ResponseWriter, r *http.Request) {
// 	var permission DistributorPermissions

// 	err := json.NewDecoder(r.Body).Decode(&permission)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprintf(w, "Invalid request payload")
// 		return
// 	}

// 	distributorID := r.Header.Get("distributor-id")
// 	distributor := distributorMap[distributorID]

// 	if distributor == nil {
// 		distributor = &Distributor{
// 			ID:          distributorID,
// 			Permissions: &permission,
// 			TrieRoot:    &TrieNode{children: make(map[string]*TrieNode)},
// 		}
// 		fmt.Printf(" permission - %#v\n", distributor.Permissions.Exclude)
// 		distributorMap[distributorID] = distributor
// 	} else {
// 		distributor.Permissions = &permission
// 		distributor.TrieRoot = &TrieNode{children: make(map[string]*TrieNode)}
// 	}

// 	addToTrie(distributor.TrieRoot, permission.Include, permission.Exclude)

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Permissions added successfully")
// }

// // Function to assign permissions from one distributor to another
// func assignPermission(w http.ResponseWriter, r *http.Request) {
// 	distributorID := r.Header.Get("distributor-id")
// 	adminID := r.Header.Get("admin-id")
// 	fmt.Println("admin id  : ", adminID)

// 	adminDistributor := distributorMap[adminID]

// 	if adminDistributor == nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprintf(w, "admin distributor not found")
// 		return
// 	}
// 	if len(distributorID) == 0 {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprintf(w, "distributor is not provided")
// 		return
// 	}

// 	var permission DistributorPermissions

// 	err := json.NewDecoder(r.Body).Decode(&permission)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprintf(w, "Invalid request payload")
// 		return
// 	}

// 	distributor := distributorMap[distributorID]
// 	if distributor == nil {
// 		distributor = &Distributor{
// 			ID:          distributorID,
// 			Permissions: &permission,
// 			TrieRoot:    &TrieNode{children: make(map[string]*TrieNode)},
// 		}
// 		distributorMap[distributorID] = distributor
// 	} else {
// 		distributor.Permissions = &permission
// 		distributor.TrieRoot = &TrieNode{children: make(map[string]*TrieNode)}
// 	}

// 	for _, d := range distributor.Permissions.Include {
// 		segments := strings.Split(d, ",")
// 		var country, state, city string
// 		if len(segments) == 1 {
// 			country = segments[0]
// 		}
// 		if len(segments) == 2 {
// 			country = segments[0]
// 			state = segments[1]
// 		}
// 		if len(segments) == 3 {
// 			country = segments[0]
// 			state = segments[1]
// 			city = segments[2]
// 		}

// 		if !hasPermission(adminDistributor.TrieRoot, adminDistributor.Permissions.Include, adminDistributor.Permissions.Exclude, country, state, city) {
// 			w.WriteHeader(http.StatusForbidden)
// 			fmt.Fprintf(w, "don't have access to permit the user with the above permission")
// 			return
// 		}
// 	}

// 	for _, d := range distributor.Permissions.Exclude {
// 		segments := strings.Split(d, ",")
// 		var country, state, city string
// 		if len(segments) == 1 {
// 			country = segments[0]
// 		}
// 		if len(segments) == 2 {
// 			country = segments[0]
// 			state = segments[1]
// 		}
// 		if len(segments) == 3 {
// 			country = segments[0]
// 			state = segments[1]
// 			city = segments[2]
// 		}

// 		if !hasPermission(adminDistributor.TrieRoot, adminDistributor.Permissions.Include, adminDistributor.Permissions.Exclude, country, state, city) {
// 			w.WriteHeader(http.StatusForbidden)
// 			fmt.Fprintf(w, "don't have access to permit the user with the above permission")
// 			return
// 		}
// 	}
// 	fmt.Printf("root permission exclude - %#v\n", adminDistributor.Permissions.Exclude)
// 	permission.Exclude = append(permission.Exclude, adminDistributor.Permissions.Exclude...)
// 	fmt.Printf("root permission final exclude - %#v\n", permission.Exclude)
// 	addToTrie(distributor.TrieRoot, permission.Include, permission.Exclude)

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Permissions assigned successfully")
// }

// // Function to check if a distributor has permission for a given input
// func checkPermission(w http.ResponseWriter, r *http.Request) {
// 	distributorID := r.Header.Get("distributor-id")

// 	distributor := distributorMap[distributorID]

// 	if distributor == nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprintf(w, "Distributor not found")
// 		return
// 	}

// 	var input struct {
// 		Country string `json:"country"`
// 		State   string `json:"state"`
// 		City    string `json:"city"`
// 	}

// 	err := json.NewDecoder(r.Body).Decode(&input)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprintf(w, "Invalid request payload")
// 		return
// 	}

// 	if hasPermission(distributor.TrieRoot, distributor.Permissions.Include, distributor.Permissions.Exclude, input.Country, input.State, input.City) {
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintf(w, "Distributor has permission")
// 	} else {
// 		w.WriteHeader(http.StatusForbidden)
// 		fmt.Fprintf(w, "Distributor does not have permission")
// 	}
// }

// // Helper function to add permissions to the Trie
// func addToTrie(root *TrieNode, include []string, exclude []string) {
// 	for _, region := range include {
// 		addRegionToTrie(root, region, true)
// 	}

// 	for _, region := range exclude {
// 		addRegionToTrie(root, region, false)
// 	}
// }

// // Helper function to add a region to the Trie
// func addRegionToTrie(node *TrieNode, region string, include bool) {
// 	// fmt.Printf("init node - %#v\n", node)
// 	segments := strings.Split(region, ",")
// 	// fmt.Println("segements : ", segments)
// 	if len(segments) == 1 {
// 		// Mark the leaf node to indicate permission
// 		if include {
// 			node.children["*"] = nil
// 		} else {
// 			node.children["-"] = nil
// 		}
// 	}

// 	for _, segment := range segments {
// 		segment = strings.TrimSpace(segment)
// 		// fmt.Println("segment : ", segment)
// 		child, ok := node.children[segment]
// 		// fmt.Printf("child node - %#v\n", child)
// 		// fmt.Println("ok - ", ok)
// 		if !ok {
// 			child = &TrieNode{children: make(map[string]*TrieNode)}

// 			// fmt.Printf("not ok child - %#v\n", child)
// 			child.children["*"] = nil
// 			node.children[segment] = child
// 			// fmt.Printf("node after not ok set - %#v\n", node)

// 		}

// 		node = child
// 		// fmt.Printf("next process node - %#v\n", node)

// 	}

// 	// Mark the leaf node to indicate permission
// 	if include {
// 		node.children["*"] = nil
// 	} else {
// 		node.children["-"] = nil
// 	}
// }

// // Helper function to check if a distributor has permission for a given input
// func hasPermission(node *TrieNode, include []string, exclude []string, country, state, city string) bool {
// 	// fmt.Printf("root - %#v\n", node)
// 	segments := []string{country, state, city}

// 	for _, segment := range segments {
// 		// fmt.Printf("node - %#v\n", node)
// 		segment = strings.TrimSpace(segment)
// 		// fmt.Println("segment : ", segment)
// 		child, ok := node.children[segment]
// 		// fmt.Printf("child - %#v\n", child)
// 		// fmt.Printf("ok - %#v\n", ok)
// 		if !ok {
// 			// Check if the leaf node indicates permission or exclusion
// 			_, okInclude := node.children["*"]
// 			_, okExclude := node.children["-"]

// 			if okInclude && !okExclude {
// 				// fmt.Println("go on")
// 				return true
// 			}
// 			return false
// 		}

// 		node = child
// 	}

// 	// Check if the leaf node indicates permission or exclusion
// 	_, okInclude := node.children["*"]
// 	_, okExclude := node.children["-"]

// 	if okInclude && !okExclude {
// 		return true
// 	}

// 	return false
// }

// func main() {
// 	http.HandleFunc("/add-permission", addPermission)
// 	http.HandleFunc("/assign-permission", assignPermission)
// 	http.HandleFunc("/checkPermission", checkPermission)

// 	log.Fatal(http.ListenAndServe(":8012", nil))
// }
