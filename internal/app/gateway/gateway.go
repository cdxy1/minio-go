package gateway

import (
	"github.com/gin-gonic/gin"

	"github.com/cdxy1/go-file-storage/internal/infra/kafka/producer"
	"github.com/cdxy1/go-file-storage/internal/repo"
	"github.com/cdxy1/go-file-storage/internal/routes/http"
	"github.com/cdxy1/go-file-storage/internal/service"
	"github.com/cdxy1/go-file-storage/internal/storage/postgres"
)

func NewApp() *gin.Engine {
	db, err := postgres.NewPostgres()

	if err != nil {
		panic("db error")
	}
	defer db.Close()

	r := gin.Default()

	fr := repo.Metadata{Db: db}
	fs := service.MetadataService{Repo: &fr}

	prod, err := producer.NewProducer()
	if err != nil {
		panic("kafka error")
	}

	http.NewMetadataHandler(r, &fs)
	http.NewFileHandler(r, prod)

	return r
}
