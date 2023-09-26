package model

import "time"

// album represents data about a record album.
type BioLink struct {
	Title     string    `json:"title"`
	Link      string    `json:"link"`
	IsPublic  bool      `json:"isPublic"`
	CreatedAt time.Time `json:"created_at" `
	UpdatedAt time.Time `json:"updated_at" `
}
