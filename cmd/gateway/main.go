package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/cdxy1/go-file-storage/internal/app/gateway"
)

func main() {
	srv := gateway.NewApp()

	srv.Static("/docs", "./api/rest")

	url := ginSwagger.URL("/docs/swagger.json")
	srv.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	srv.Run(":8080")
}
