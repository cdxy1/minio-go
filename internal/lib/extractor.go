package lib

import (
	"encoding/json"
	"mime/multipart"
	"time"

	"github.com/google/uuid"

	"github.com/cdxy1/go-file-storage/internal/entity"
)

func ExtractMetadata(fileHeader *multipart.FileHeader) ([]byte, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	metadata := &entity.Metadata{
		Id:        newUUID.String(),
		Name:      fileHeader.Filename,
		Url:       "placeholder",
		Size:      fileHeader.Size,
		Type:      fileHeader.Header.Get("Content-Type"),
		CreatedAt: time.Now(),
	}

	msg, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
