package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/cdxy1/go-file-storage/internal/entity"
	"github.com/cdxy1/go-file-storage/internal/service"
)

type FileHandler struct {
	fs *service.FileService
}

func NewFileHandler(r *gin.Engine, fs *service.FileService) *FileHandler {
	fh := &FileHandler{fs}

	files := r.Group("/file")
	{
		files.POST("", fh.Create)
		files.GET(":id", fh.Find)
	}
	return fh
}

func (fh *FileHandler) Create(c *gin.Context) {
	var f entity.File

	err := c.ShouldBindJSON(&f)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := fh.fs.CreateFile(c, &f); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, f)
}

func (fh *FileHandler) Find(c *gin.Context) {
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
