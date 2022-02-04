package core

import (
	"sort"

	"github.com/finleygn/soundcloud-watch/pkg/core/models"
)

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

	sort.Ints(ids)
	return ids
}

// Find added and remove tracks between two lists.
// Lists must be pre sorted.
func Diff(prev []int, new []int) (added []int, removed []int) {
	for _, item := range prev {
		if sort.SearchInts(new, item) != len(new) {
			removed = append(removed, item)
		}
	}

	for _, item := range new {
		if sort.SearchInts(prev, item) != len(new) {
			removed = append(added, item)
		}
	}

	return added, removed
}
