package main

import (
	"fmt"
	"os"

	"github.com/jeffmccune/tracks"
)

func main() {
	tracks, err := tracks.GetTrackInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	var start float64

	for _, track := range tracks {
		fmt.Printf("%.6f\t%.6f\t%s\n", start, start, track.Name)
		start += track.Duration
	}

	os.Exit(0)
}
