package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/finleygn/soundcloud-watch/pkg/client"
	"github.com/finleygn/soundcloud-watch/pkg/core"
	"github.com/finleygn/soundcloud-watch/pkg/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "scwatch",
	Short:         "Watch for changes in soundcloud likes.",
	SilenceErrors: false,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Check users latest likes",
	RunE: func(cmd *cobra.Command, args []string) error {
		username := cmd.Flag("user").Value.String()

		userDir := core.OpenUserDir(username)

		if !userDir.HasLatestState() {
			return errors.New("user is not initialized")
		}

		user, err := client.GetUser("atkWGyMg57QFFAwK5c9VpC1N5Q141g7I", username)
		if err != nil {
			return err
		}

		_, chunk := ui.CreateProgressBar(user.LikesTotal)
		likes, err := user.GetAllLikes(50, chunk)
		if err != nil {
			return err
		}

		fmt.Println()

		known, err := userDir.ReadKnownTracks()
		if err != nil {
			return err
		}

		last_state, err := userDir.ReadLatestState()
		if err != nil {
			panic(err)
		}

		current_likes := core.TracksToIds(likes)

		added := core.FindAdded(last_state.All, current_likes)
		removed := core.FindRemoved(last_state.All, current_likes)

		new_state := core.State{
			All:     current_likes,
			Added:   added,
			Removed: removed,
		}

		userDir.WriteLatestState(new_state)
		ui.InfoMessage("Saved latest state")

		backup_err := userDir.WriteBackup(new_state)
		if backup_err != nil {
			return backup_err
		}
		ui.InfoMessage("Created backup of latest state")

		// Update now known tracks
		for _, track := range likes {
			known[track.Id] = track
		}
		userDir.WriteKnownTracks(known)

		ui.InfoMessage("Updated known tracks")

		ui.SuccessMessage("Complete.")

		ui.PrintState(new_state, known)

		return nil
	},
}

var initUserCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a log for user likes",
	RunE: func(cmd *cobra.Command, args []string) error {
		username := cmd.Flag("user").Value.String()
		userDir := core.OpenUserDir(username)

		if userDir.HasLatestState() {
			cmd.SilenceUsage = true
			return errors.New("this user is already initialized")
		}

		user, err := client.GetUser("atkWGyMg57QFFAwK5c9VpC1N5Q141g7I", username)
		if err != nil {
			return err
		}

		_, cb := ui.CreateProgressBar(user.LikesTotal)
		likes, err := user.GetAllLikes(50, cb)
		if err != nil {
			return err
		}

		// // Save default latest state
		userDir.WriteLatestState(core.State{
			All:     core.TracksToIds(likes),
			Added:   make([]int, 0),
			Removed: make([]int, 0),
		})

		fmt.Println("Created initial state for", username)

		return nil
	},
}

var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "Print latest run stats",
	RunE: func(cmd *cobra.Command, args []string) error {
		userDir := core.OpenUserDir(cmd.Flag("user").Value.String())

		known, err := userDir.ReadKnownTracks()
		if err != nil {
			return err
		}

		last_state, err := userDir.ReadLatestState()
		if err != nil {
			return err
		}

		ui.PrintState(*last_state, known)

		return nil
	},
}

var songCmd = &cobra.Command{
	Use:   "song",
	Short: "Print details of a song",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		userDir := core.OpenUserDir(cmd.Flag("user").Value.String())

		known, err := userDir.ReadKnownTracks()
		if err != nil {
			return err
		}

		track, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		ui.PrintTrack(known[track])

		return nil
	},
}

func main() {
	pf := rootCmd.PersistentFlags()

	pf.StringP("user", "u", "", "username")
	cobra.MarkFlagRequired(pf, "user")

	rootCmd.AddCommand(statCmd)
	rootCmd.AddCommand(initUserCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(songCmd)

	rootCmd.DisableSuggestions = true

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
