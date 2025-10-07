package entity

import "time"

type File struct {
	Id        int
	Name      string
	Url       string
	CreatedAt time.Time
}
