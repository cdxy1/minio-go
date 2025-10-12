package main

import (
	"github.com/cdxy1/go-file-storage/internal/app/gateway"
)

func main() {
	srv := gateway.NewApp()
	srv.Run()
}
