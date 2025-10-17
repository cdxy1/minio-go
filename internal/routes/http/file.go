package http

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cdxy1/go-file-storage/internal/grpc/file"
	grpcclient "github.com/cdxy1/go-file-storage/internal/infra/grpc_client"
)

func NewFileHandler(r *gin.Engine) {
	client, err := grpcclient.NewFileGrpcClient()
	if err != nil {
		println(err.Error())
		panic("blabla")
	}

	file := r.Group("/file")
	{
		file.GET(":id/download", func(c *gin.Context) {
			Download(c, client)
		})
		file.POST("upload", func(c *gin.Context) {
			Upload(c, client)
		})
	}

}

func Download(c *gin.Context, client file.FileServiceClient) {
	fileId := c.Param("id")

	resp, err := client.DownloadFile(c, &file.DownloadFileRequest{Name: fileId})

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	data := resp.Data

	c.Data(http.StatusOK, "application/octet-stream", data)
}

func Upload(c *gin.Context, client file.FileServiceClient) {

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileData, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer fileData.Close()

	data, err := io.ReadAll(fileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := client.UploadFile(c, &file.UploadFileRequest{Name: fileHeader.Filename, Data: data})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Name)
}
