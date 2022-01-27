package client

import (
	"errors"
	"fmt"

	"github.com/finleygn/soundcloud-watch/pkg/client/models"
	"github.com/finleygn/soundcloud-watch/pkg/client/parser"
	coremodel "github.com/finleygn/soundcloud-watch/pkg/core/models"
)

type User struct {
	id         int
	username   string
	clientId   string
	LikesTotal int
}

func GetUser(clientId string, username string) (user *User, err error) {
	hydration_data, err := GetHydrationData(username)

	if err != nil {
		return nil, err
	}

	user = &User{
		id:         hydration_data.Id,
		username:   username,
		clientId:   clientId,
		LikesTotal: hydration_data.Likes,
	}

	user.id = hydration_data.Id
	user.LikesTotal = hydration_data.Likes

	return user, nil
}

func (u *User) GetLikesFromOffset(offset string, limit int) (response *models.LikeResponse, err error) {
	body, resp, err := get(
		fmt.Sprintf(
			"https://api-v2.soundcloud.com/users/%d/likes?client_id=%s&limit=%d&offset=%s&app_locale=en",
			u.id,
			u.clientId,
			limit,
			offset,
		),
	)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("unexpected code from SoundCloud")
	}

	res := parser.ParseLikeResponse(body)
	return &res, nil
}

func (u *User) GetAllLikes(limit int, on_chunk func()) (likes []coremodel.Track, err error) {
	likes = []coremodel.Track{}
	offset := "0"

	for {
		response, err := u.GetLikesFromOffset(offset, limit)
		on_chunk()

		if err != nil {
			return nil, err
		}

		for _, item := range response.Collection {
			if item.Track.Id != 0 {
				likes = append(likes, item.Track)
			}
		}

		offset = response.NextId

		if offset == "" {
			break
		}
	}

	return likes, nil
}
