package gateway

import (
	"github.com/gin-gonic/gin"

	"github.com/cdxy1/minio-go/internal/routes/http"
)

func NewApp() *gin.Engine {
	r := gin.Default()

	http.NewMetadataHandler(r)
	http.NewFileHandler(r)

	return r
}
