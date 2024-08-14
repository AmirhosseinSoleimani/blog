package main

import (
	"blog/config"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	configs := configSet()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"App Name": viper.Get("App.Name"),
		})
	})
	r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
}

func configSet() config.Config{
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {

		fmt.Println("Errror reading the configs")
	}
	
	var configs config.Config

	err := viper.Unmarshal(&configs)

	if err!= nil {
        fmt.Printf("unable to decode into struct, %v", err)
    }
	return configs
}
