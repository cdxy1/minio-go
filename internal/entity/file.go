package entity

import "time"

type File struct {
	Id        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
