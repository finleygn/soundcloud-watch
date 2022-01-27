package core

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"strings"
	"time"

	"github.com/finleygn/soundcloud-watch/pkg/core/models"
)

type userDir string

func (u *userDir) getBasePath() string {
	cfg_dir, err := os.UserConfigDir()

	if err != nil {
		panic("ruhoh")
	}

	return path.Join(cfg_dir, "scwatch", string(*u))
}

func (u *userDir) getKnownTracksFilePath() string {
	return path.Join(u.getBasePath(), "known-tracks.json")
}

func (u *userDir) getLatestStateFilePath() string {
	return path.Join(u.getBasePath(), "latest-state.json")
}

func (u *userDir) getBackupsPath() string {
	return path.Join(u.getBasePath(), "backup")
}

func (u *userDir) getBackupPath(time time.Time) string {
	timestring := time.Format("2006-01-02 15:04:05")
	timestring = strings.Replace(timestring, " ", "_", 1)
	timestring = timestring + ".json"

	return path.Join(u.getBackupsPath(), timestring)
}

//

func OpenUserDir(name string) userDir {
	dir := userDir(name)

	if _, err := os.Stat(dir.getBackupsPath()); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(dir.getBackupsPath(), os.ModePerm)
	}

	if _, err := os.Stat(dir.getKnownTracksFilePath()); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(dir.getBasePath(), os.ModePerm)
		os.WriteFile(dir.getKnownTracksFilePath(), []byte("{}"), os.ModePerm)
	}

	return dir
}

//

type KnownTracks map[int]models.Track

func (u *userDir) ReadKnownTracks() (KnownTracks, error) {
	file, err := os.ReadFile(u.getKnownTracksFilePath())

	if err != nil {
		return nil, err
	}

	var knownTracks KnownTracks
	json.Unmarshal(file, &knownTracks)

	if err := json.Unmarshal(file, &knownTracks); err != nil {
		return nil, err
	}

	return knownTracks, nil
}

func (u *userDir) WriteKnownTracks(value KnownTracks) error {
	json, err := json.Marshal(value)

	if err != nil {
		return err
	}

	write_error := os.WriteFile(u.getKnownTracksFilePath(), json, 0644)

	if write_error != nil {
		return write_error
	}

	return nil
}

func (u *userDir) ReadLatestState() (*State, error) {
	file, err := os.ReadFile(u.getLatestStateFilePath())

	if err != nil {
		return nil, err
	}

	var state State
	json.Unmarshal(file, &state)

	if err := json.Unmarshal(file, &state); err != nil {
		return nil, err
	}

	return &state, nil
}

func (u *userDir) WriteLatestState(value State) error {
	json, err := json.Marshal(value)

	if err != nil {
		return err
	}

	write_error := os.WriteFile(u.getLatestStateFilePath(), json, 0644)

	if write_error != nil {
		return write_error
	}

	return nil
}

func (u *userDir) HasLatestState() bool {
	if _, err := os.Stat(u.getLatestStateFilePath()); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func (u *userDir) WriteBackup(value State) error {
	json, err := json.Marshal(value)

	if err != nil {
		return err
	}

	write_error := os.WriteFile(u.getBackupPath(time.Now()), json, 0644)

	if write_error != nil {
		return write_error
	}

	return nil
}
