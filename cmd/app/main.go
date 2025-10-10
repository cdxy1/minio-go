package main

import (
	"github.com/cdxy1/go-file-storage/internal/app"
)

func main() {
	srv := app.NewApp()
	srv.Run()
}
