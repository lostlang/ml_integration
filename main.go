package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "ml_integration/docs"
	"ml_integration/handler"
)

// @title ML integration course
// @version 0.0.2
// @description Реализация домашних заданий для курса
// @host localhost:7577
// @BasePath /
func main() {
	r := gin.Default()

	r.POST("calculate", handler.SumHandler)
	r.POST("query", handler.QueryHandler)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":7577")
}
