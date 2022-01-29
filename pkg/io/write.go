package io

import (
	"fmt"
	"io"

	gpxgo "github.com/tkrajina/gpxgo/gpx"
)

func Write(w io.Writer, gpx gpxgo.GPX) error {
	outData, err := gpx.ToXml(gpxgo.ToXmlParams{Version: "1.1", Indent: true})
	if err != nil {
		return fmt.Errorf("converting rotated content to XML: %w", err)
	}

	if _, err := w.Write(outData); err != nil {
		return fmt.Errorf("writing rotated GPX data: %w", err)
	}
	return nil
}
