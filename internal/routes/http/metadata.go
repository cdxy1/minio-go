package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-gonic/gin"

	"github.com/cdxy1/go-file-storage/internal/entity"
	"github.com/cdxy1/go-file-storage/internal/service"
)

type MetadataHandler struct {
	fs *service.MetadataService
}

func NewMetadataHandler(r *gin.Engine, fs *service.MetadataService) *MetadataHandler {
	md := &MetadataHandler{fs}

	metadata := r.Group("/metadata")
	{
		metadata.GET(":id", md.Find)
	}
	return md
}

func (fh *MetadataHandler) Find(c *gin.Context) {
	idPath := c.Param("id")
	id, err := strconv.Atoi(idPath)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid param"})
		return
	}

	res, err := fh.fs.GetFile(c, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (fh *MetadataHandler) HandleMessage(msg []byte, offset kafka.Offset) error {
	data := entity.Metadata{}
	if err := json.Unmarshal(msg, &data); err != nil {
		println("fuyegwiugyfewuhygewfefwuihyg")
		return err
	}
	fmt.Println(data)
	// fh.fs.CreateFile(context.Background(), )
	return nil
}
