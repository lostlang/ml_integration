package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "ml_integration/docs"
	"ml_integration/handler"
)

// @title ML integration course
// @version 0.0.3
// @description Реализация домашних заданий для курса
// @BasePath /
func main() {
	r := gin.Default()

	r.POST("calculate", handler.SumHandler)
	r.POST("query", handler.QueryHandler)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(":7577")
	if err != nil {
		fmt.Println(err)
		return
	}
}
