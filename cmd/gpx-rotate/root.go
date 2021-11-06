package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/glynternet/gpx-rotate/pkg/gpxrotate"
	"github.com/glynternet/pkg/log"
	"github.com/spf13/cobra"
	gpxgo "github.com/tkrajina/gpxgo/gpx"
)

func buildCmdTree(logger log.Logger, out io.Writer, rootCmd *cobra.Command) {
	var rotation int
	*rootCmd = cobra.Command{
		Use:   "gpx-rotate <gpx-file> <points>",
		Short: "Rotate a single track GPX file by a given number of points.",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.ExactArgs(2)(cmd, args); err != nil {
				return err
			}
			validRotation, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("points arg must be integer: %w", err)
			}
			rotation = validRotation
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			gpxData, err := ioutil.ReadFile(path)
			if err != nil {
				return fmt.Errorf("reading gpx file:%q: %w", path, err)
			}

			gpx, err := gpxgo.ParseBytes(gpxData)
			if err != nil {
				return fmt.Errorf("parsing gpx data: %w", err)
			}

			if err := validate(gpx); err != nil {
				return fmt.Errorf("gpx invalid or not supported: %w", err)
			}

			gpx.Tracks[0].Segments[0].Points = gpxrotate.Rotated(gpx.Tracks[0].Segments[0].Points, rotation)

			outData, err := gpx.ToXml(gpxgo.ToXmlParams{Version: "1.1", Indent: true})
			if err != nil {
				return fmt.Errorf("converting rotated gpx to XML: %w", err)
			}

			if _, err := out.Write(outData); err != nil {
				return fmt.Errorf("writing rotated GPX data: %w", err)
			}
			return nil
		},
	}
}

func validate(g *gpxgo.GPX) error {
	if n := len(g.Tracks); n != 1 {
		return fmt.Errorf("gpx file must contain exactly 1 track but contains %d", n)
	}

	if n := len(g.Tracks[0].Segments); n != 1 {
		return fmt.Errorf("track must contain exactly 1 segment but contains %d", n)
	}

	gpxrotate.Rotated(g.Tracks[0].Segments[0].Points, 10)
	return nil
}
