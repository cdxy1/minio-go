package main

import (
	"github.com/gin-gonic/gin"

	"github.com/cdxy1/go-file-storage/internal/database"
	"github.com/cdxy1/go-file-storage/internal/repo"
	"github.com/cdxy1/go-file-storage/internal/routes"
	"github.com/cdxy1/go-file-storage/internal/service"
)

func main() {
	db, err := database.NewPostgres()

	if err != nil {
		panic("db error")
	}
	defer db.Close()

	r := gin.Default()

	fr := repo.File{Db: db}
	fs := service.FileService{Repo: &fr}

	routes.NewFileHandler(r, &fs)

	r.Run(":8080")
}
