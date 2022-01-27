package parser

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/finleygn/soundcloud-watch/pkg/client/models"
)

func ParseLikeResponse(body []byte) models.LikeResponse {
	var data models.LikeResponse

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Print(err)
		panic("Soundcloud changed their like API... This program doesn't work anymore")
	}

	if data.NextHref != "" {
		regex, _ := regexp.Compile(`\?offset=(.+)&`)
		next_id := regex.FindStringSubmatch(data.NextHref)[1]

		data.NextId = next_id
	}

	return data
}
