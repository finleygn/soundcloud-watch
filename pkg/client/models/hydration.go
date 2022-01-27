package models

type HydrationResponse struct {
	Id    int
	Likes int
}

type HydrationResponseField struct {
	Id         int `json:"id"`
	LikesCount int `json:"likes_count"`
}
