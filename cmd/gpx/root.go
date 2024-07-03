package main

import (
	"io"

	"github.com/glynternet/pkg/log"
	"github.com/spf13/cobra"
)

func buildCmdTree(logger log.Logger, out io.Writer, rootCmd *cobra.Command) {
	rootCmd.AddCommand(csvCmd(out))
	rootCmd.AddCommand(jsonCmd(logger, out))
	rootCmd.AddCommand(rotateCmd(out))
	rootCmd.AddCommand(splitFileCmd(logger))
	rootCmd.AddCommand(splitTrackCmd(logger))
	rootCmd.AddCommand(trackWaypointsCmd(logger, out))
}
