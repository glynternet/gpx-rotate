package main

import (
	gpxhttp "github.com/glynternet/gpx/internal/http"
	"io"
	"net/http"

	"github.com/glynternet/pkg/log"
	"github.com/spf13/cobra"
)

func buildCmdTree(logger log.Logger, out io.Writer, rootCmd *cobra.Command) {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "serve",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return (&http.Server{
				Addr:    ":8080",
				Handler: http.HandlerFunc(gpxhttp.HandleElevation),
			}).ListenAndServe()
		},
	})
}
