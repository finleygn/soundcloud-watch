package models

import "github.com/finleygn/soundcloud-watch/pkg/core/models"

type CollectionItemResponse struct {
	CreatedAt string       `json:"created_at"`
	Kind      string       `json:"kind"`
	Track     models.Track `json:"track"`
}

type LikeResponse struct {
	Collection []CollectionItemResponse `json:"collection"`
	NextHref   string                   `json:"next_href"`
	NextId     string                   `json:"next_id"`
}
