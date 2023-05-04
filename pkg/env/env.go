package env

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	AppPort        string `split_words:"true"`
	CitiesFileName string `split_words:"true"`
}

var EnvCfg EnvConfig

func ProcssEnv() error {
	if err := envconfig.Process("", &EnvCfg); err != nil {
		return err
	}
	return nil
}
func InitEnv() {
	configFile := "pkg/env/.env"
	if err := godotenv.Load(configFile); err != nil {
		log.Fatal("unable to load .env")
	}
	if err := ProcssEnv(); err != nil {
		log.Fatal("unable to load .env")
	}
	cfgStr, _ := json.Marshal(EnvCfg)
	log.Println(string(cfgStr))
}
