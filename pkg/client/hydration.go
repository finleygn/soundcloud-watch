package client

import (
	"errors"
	"fmt"

	"github.com/finleygn/soundcloud-watch/pkg/client/models"
	"github.com/finleygn/soundcloud-watch/pkg/client/parser"
)

func GetHydrationData(username string) (response *models.HydrationResponse, err error) {
	body, resp, err := get(fmt.Sprintf("https://soundcloud.com/%s", username))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 404 {
		return nil, errors.New("username does not exist")
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("unexpected code from SoundCloud")
	}

	data := parser.ParseHydrationResponse(body)

	return &data, nil
}
