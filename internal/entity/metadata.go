package entity

import "time"

type Metadata struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	Size      int64     `json:"size"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
