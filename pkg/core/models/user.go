package models

type User struct {
	AvatarUrl    string `json:"avatar_url"`
	Id           int    `json:"id"`
	LastModified string `json:"last_modified"`
	Permalink    string `json:"permalink"`
	Uri          string `json:"uri"`
	Urn          string `json:"urn"`
	Username     string `json:"username"`
}
