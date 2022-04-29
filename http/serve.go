package http

import (
	"fmt"
	"net/http"
	"time"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// ServeStart init the serve main function
func ServeStart() {
	viper.SetConfigFile("config.json")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error in config file: %s ", err))
	}

	gin.SetMode(viper.GetString("HTTP.enviroment")) // release or debug

	router := gin.Default()

	initGinConfig(router)
	Router(router)
	serve(router)
}

func initGinConfig(router *gin.Engine) {

	router.Use(limit.MaxAllowed(200))
	router.Use(cors.Default())
}

func serve(router *gin.Engine) {
	viper.SetConfigFile("config.json")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error in config file: %s ", err))
	}

	serverPort := fmt.Sprintf(":%s", viper.GetString("HTTP.port"))
	s := &http.Server{
		Addr:           serverPort,
		Handler:        router,
		ReadTimeout:    18000 * time.Second,
		WriteTimeout:   18000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		panic(fmt.Errorf("Fatal Error Description: %s ", err))
	}
}
