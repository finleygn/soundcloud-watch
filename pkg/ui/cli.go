package ui

import (
	"fmt"

	"github.com/finleygn/soundcloud-watch/pkg/core"
	"github.com/finleygn/soundcloud-watch/pkg/core/models"
	"github.com/mitchellh/colorstring"
	"github.com/schollz/progressbar/v3"
)

func CreateProgressBar(total int) (*progressbar.ProgressBar, func()) {
	bar := progressbar.NewOptions(
		int(float64(total)/float64(50)),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetDescription("[cyan]Fetching likes[reset]"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[cyan]=[reset]",
			SaucerHead:    "[cyan]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)
	return bar, func() { bar.Add(1) }
}

func SuccessMessage(message string) {
	colorstring.Println("[green]✅ " + message + "[reset]")
}

func InfoMessage(message string) {
	colorstring.Println("[cyan]ℹ️  " + message + "[reset]")
}

func ErrorMessage(message string) {
	colorstring.Println("[red]❌ " + message + "[reset]")
}

func squareBracket(content string) string {
	return colorstring.Color("[dark_gray][[reset]") + content + colorstring.Color("[dark_gray]][reset]")
}

func PrintState(state core.State, known core.KnownTracks) {
	var out string

	out += "\n"
	out += colorstring.Color("[green][+ Added][reset]\n")

	if len(state.Added) > 0 {
		for _, id := range state.Added {
			track := known[id]
			id_str := fmt.Sprintf("%10d", track.Id)

			// Some tracks are equal to 0
			if id == 0 {
				continue
			}

			out += fmt.Sprintf(
				"%s%s %s\n",
				squareBracket(colorstring.Color("[cyan]"+id_str)),
				squareBracket(colorstring.Color("[yellow]"+track.User.Username)),
				track.Title,
			)
		}
	} else {
		out += colorstring.Color("[dark_gray]No updates[reset]\n")
	}

	out += "\n"
	out += colorstring.Color("[red][- Removed][reset]\n")

	if len(state.Removed) > 0 {
		for _, id := range state.Removed {
			track := known[id]
			id_str := fmt.Sprintf("%10d", track.Id)

			// Some tracks are equal to 0
			if id == 0 {
				continue
			}

			out += fmt.Sprintf(
				"%s%s %s\n",
				squareBracket(colorstring.Color("[cyan]"+id_str)),
				squareBracket(colorstring.Color("[yellow]"+track.User.Username)),
				track.Title,
			)
		}
	} else {
		out += colorstring.Color("[dark_gray]No updates[reset]\n")
	}

	fmt.Println(out)
}

func PrintTrack(track models.Track) {
	// TODO:
	fmt.Println(track)
}
