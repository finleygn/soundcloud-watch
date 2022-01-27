package core

import "github.com/finleygn/soundcloud-watch/pkg/core/models"

type State struct {
	All     []int `json:"all"`
	Added   []int `json:"added"`
	Removed []int `json:"removed"`
}

func TracksToIds(tracks []models.Track) []int {
	ids := []int{}

	for _, track := range tracks {
		ids = append(ids, track.Id)
	}

	return ids
}

func contains(items []int, item int) bool {
	for _, current := range items {
		if current == item {
			return true
		}
	}
	return false
}

// TODO: Needs massive optimisation lmao i just wanna see it working

func FindRemoved(prev []int, new []int) []int {
	removed := []int{}

	for _, item := range prev {
		if !contains(new, item) {
			removed = append(removed, item)
		}
	}

	return removed
}

func FindAdded(prev []int, new []int) []int {
	added := []int{}

	for _, item := range new {
		if !contains(prev, item) {
			added = append(added, item)
		}
	}

	return added
}
