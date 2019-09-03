package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/log"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/mongo"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/server"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "c", "", "config filename")
}

func initConfig() {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println("using config file", viper.ConfigFileUsed())
}

func initLog() {
	log.Init()
}

func initMongo() {
	var logisticConfig mongo.MongoConfig
	if err := viper.UnmarshalKey("mongo", &logisticConfig); err != nil {
		log.Fatal(errors.Wrap(err, "read mongo config"))
	}

	// init mongo db
	if err := mongo.InitMongo(&logisticConfig); err != nil {
		log.Fatal(errors.Wrap(err, "init mongo"))
	}
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()

	v1 := router.Group("logistic/v1")
	kuaidi100 := v1.Group("/kuaidi100")
	{
		kuaidi100.POST("/callback", server.K1Callback)
	}

	logistic := v1.Group("/logistic")
	{
		logistic.GET("/query", server.LogisticQuery)
		logistic.GET("/permission", server.PermissionQuery)
	}

	return router
}

func main() {
	// INIT
	flag.Parse()
	if configFile == "" {
		panic("must specified config file")
	}

	initConfig()
	initLog()
	initMongo()

	r := setupRouter()
	r.Run(":" + viper.GetString("listen"))
}
