package tracks

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Track struct {
	Duration float64
	Name     string
}

func GetTrackInfo() ([]Track, error) {
	script := `tell application "Music"
        set trackList to ""
        repeat with t in selection
            set trackList to trackList & (duration of t) & tab & (name of t) & linefeed
        end repeat
        return trackList
    end tell`

	cmd := exec.Command("osascript", "-e", script)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var tracks []Track
	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "\t")
		if len(parts) != 2 {
			continue
		}

		duration, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse duration: %v", err)
		}

		tracks = append(tracks, Track{
			Duration: duration,
			Name:     strings.TrimSpace(parts[1]),
		})
	}

	return tracks, nil
}
