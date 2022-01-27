package models

type Track struct {
	CreatedAt    string `json:"created_at"`
	Description  string `json:"description"`
	Id           int    `json:"id"`
	LikesCount   int    `json:"likes_count"`
	Title        string `json:"title"`
	PermalinkUrl string `json:"permalink_url"`
	RepostsCount int    `json:"reposts_count"`
	ArtworkUrl   string `json:"artwork_url"`
	User         User   `json:"user"`
}
