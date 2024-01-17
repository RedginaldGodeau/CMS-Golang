package main

import (
	"cms/internal/config"
	"cms/internal/migration"
	"cms/internal/transports/api"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		println(err.Error())
		return
	}

	err = migration.Start()
	if err != nil {
		println("Migration: ", err.Error())
	} else {
		println("Migration: Done...")
	}

	r := gin.Default()
	r.LoadHTMLGlob("./public/templates/*/**")

	apiRoute := r.Group("v1/api")
	{
		api.Init(apiRoute)
	}

	err = r.Run(":" + viper.GetString("server.port"))
	if err != nil {
		println(err.Error())
		return
	}
}
