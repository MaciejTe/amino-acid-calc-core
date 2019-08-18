package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/toorop/gin-logrus"
	"io"
	"os"
)

func SetupRouter() *gin.Engine {
	// create router with all necessary properties
	router := gin.Default()
	//router.Use(gin.Logger())
	log := logrus.New()
	router.Use(ginlogrus.Logger(log), gin.Recovery())
	logFile, _ := os.Create("gin.log")
	mw := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(mw)

	// write to the log file and stdout at the same time
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	// v1 API group
	v1 := router.Group("/api/v1")
	{
		EndpointGroup := v1.Group("/")
		{
			EndpointGroup.GET("/sample", Sample)
		}
	}

	return router
}
