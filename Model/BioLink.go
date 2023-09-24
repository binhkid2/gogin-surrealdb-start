package model

// album represents data about a record album.
type BioLink struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	IsPublic bool   `json:"isPublic"`
}
