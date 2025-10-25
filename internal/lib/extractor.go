package lib

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/cdxy1/go-file-storage/internal/entity"
)

func ExtractMetadata(name string, fileData []byte) ([]byte, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	metadata := &entity.Metadata{
		Id:        newUUID.String(),
		Name:      name,
		Url:       "placeholder",
		Size:      int64(len(fileData)),
		Type:      http.DetectContentType(fileData),
		CreatedAt: time.Now(),
	}

	msg, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
