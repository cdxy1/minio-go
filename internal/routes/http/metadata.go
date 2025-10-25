package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

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


