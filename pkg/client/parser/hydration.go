package parser

import (
	"encoding/json"
	"regexp"

	"github.com/finleygn/soundcloud-watch/pkg/client/models"
)

func ParseHydrationResponse(body []byte) models.HydrationResponse {
	regex, _ := regexp.Compile(`window\.__sc_hydration = \[.+\{"hydratable":"user","data":(.+)}`)
	raw_data := regex.FindSubmatch(body)[1]

	var data models.HydrationResponseField

	if err := json.Unmarshal(raw_data, &data); err != nil {
		panic("Soundcloud changed their hydratable API... This program doesn't work anymore")
	}

	return models.HydrationResponse{
		Likes: data.LikesCount,
		Id:    data.Id,
	}
}
