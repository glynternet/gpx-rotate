// Code generated by dubplate v0.10.1 DO NOT EDIT.
// Implement the following function to use this boilerplate
// func buildCmdTree(logger log.Logger, out io.Writer, rootCmd *cobra.Command) {}

package main

import (
	"os"

	"github.com/glynternet/pkg/cmd"
	"github.com/glynternet/pkg/log"
	"github.com/spf13/cobra"
)

const appName = "gpx"

// to be changed using ldflags with the go build command
var version = "unknown"

func main() {
	logger := log.NewLogger(os.Stderr)
	out := os.Stdout

	var rootCmd = &cobra.Command{
		Use: appName,
	}

	rootCmd.AddCommand(cmd.NewVersionCmd(version, out))
	buildCmdTree(logger, out, rootCmd)

	if err := rootCmd.Execute(); err != nil {
		_ = logger.Log(
			log.Message("Error executing root command"),
			log.ErrorMessage(err))
		os.Exit(1)
	}
}

