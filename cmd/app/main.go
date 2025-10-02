package main

import (
	"github.com/cdxy1/go-file-storage/pkg/logger"
)

func main() {
	log := logger.SetupLogger()
	log.Debug("мамку ебал")
	log.Info("ну тут инфо")
	log.Warn("Ну тут ворнинги хехе")
}
