package main

import (
	"fmt"
	"io"

	"github.com/glynternet/pkg/log"
	"github.com/spf13/cobra"
	gpxgo "github.com/tkrajina/gpxgo/gpx"
)

func buildCmdTree(logger log.Logger, out io.Writer, rootCmd *cobra.Command) {
	rootCmd.AddCommand(rotateCmd(out))
}

func validate(g *gpxgo.GPX) error {
	if n := len(g.Tracks); n != 1 {
		return fmt.Errorf("gpx file must contain exactly 1 track but contains %d", n)
	}

	if n := len(g.Tracks[0].Segments); n != 1 {
		return fmt.Errorf("track must contain exactly 1 segment but contains %d", n)
	}

	return nil
}
