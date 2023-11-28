package config

import (
	"strings"

	"github.com/spf13/viper"

	"fmt"
)

var appConfig config

type config struct {
	appPort         int
	rpcUrl          string
	mongoDbName     string
	mongoDbUsername string
	mongoDbPassword string
	allowedOrigins  []string
	devMode         bool `default:"false"`
	storageHostname string
}

func Load() error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	appConfig = config{
		appPort:         viper.GetInt("PORT"),
		rpcUrl:          viper.GetString("RPC_URL"),
		mongoDbName:     viper.GetString("MONGODB_HOSTNAME"),
		mongoDbUsername: viper.GetString("MONGODB_USERNAME"),
		mongoDbPassword: viper.GetString("MONGODB_PASSWORD"),
		allowedOrigins:  strings.Split(viper.GetString("ALLOWED_ORIGINS"), ","),
		devMode:         viper.GetBool("DEV_MODE"),
		storageHostname: viper.GetString("STORAGE_HOSTNAME"),
	}

	return nil
}

func MongoDbConnectionUrl() string {
	srv := "+srv"

	if appConfig.devMode {
		srv = ""
	}

	return fmt.Sprintf("mongodb%s://%s:%s@%s/?retryWrites=true&w=majority", srv, appConfig.mongoDbUsername, appConfig.mongoDbPassword, appConfig.mongoDbName)
}

func AllowedOrigins() []string {
	return appConfig.allowedOrigins
}

func StorageHostname() string {
	return appConfig.storageHostname
}

func RpcUrl() string {
	return appConfig.rpcUrl
}

func DevMode() bool {
	return appConfig.devMode
}

func ServerPort() int {
	return appConfig.appPort
}
