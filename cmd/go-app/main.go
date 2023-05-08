package main

import (
	"distribution-mgmnt/app"
	"distribution-mgmnt/internal/api"
	"distribution-mgmnt/pkg/cmaps"
	"distribution-mgmnt/pkg/env"
	"distribution-mgmnt/pkg/files"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	env.InitEnv()
	cmaps.DistributorMgmntDB = cmaps.NewDistributionMaps()
	files.CSVLoader(env.EnvCfg.CitiesFileName, 1)
	router := gin.Default()
	router = app.RegisterHandlers(router, api.NewDistributionMgnmtServer())
	log.Infoln("Hey Guys Server Is Up On Running At Port: ", env.EnvCfg.AppPort)
	router.Run(":" + env.EnvCfg.AppPort)
}
