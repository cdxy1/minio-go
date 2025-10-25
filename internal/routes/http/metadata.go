package http

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/cdxy1/go-file-storage/internal/grpc/metadata"
	grpcclient "github.com/cdxy1/go-file-storage/internal/infra/grpc_client"
)

func NewMetadataHandler(r *gin.Engine) {
	client, err := grpcclient.NewMetadataGprcClient()
	if err != nil {
		panic("blabla")
	}

	v1 := r.Group("/api/v1")
	{
		metadata := v1.Group("/metadata")
		{
			metadata.GET(":id", func(ctx *gin.Context) {
				FindById(ctx, client)
			})
			metadata.GET("", func(ctx *gin.Context) {
				GetAll(ctx, client)
			})
		}
	}
}

func FindById(c *gin.Context, client metadata.MetadataServiceClient) {
	idPath := c.Param("id")

	res, err := client.GetById(c, &metadata.FileMetadataRequest{Id: idPath})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	data, err := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var obj map[string]interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, obj)
}

func GetAll(c *gin.Context, client metadata.MetadataServiceClient) {
	res, err := client.GetAll(c, &emptypb.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data, err := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var obj map[string]interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, obj)
}
