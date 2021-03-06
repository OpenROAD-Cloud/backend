package main

import (
	"github.com/OpenROAD-Cloud/backend/runner"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	runner.InitializeRoutes(router)

	router.Run()
}
