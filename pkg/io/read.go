package io

import (
	"fmt"
	"github.com/tkrajina/gpxgo/gpx"
	"io/ioutil"
)

func ReadFile(path string) (*gpx.GPX, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading content file:%q: %w", path, err)
	}

	data, err := gpx.ParseBytes(content)
	if err != nil {
		return nil, fmt.Errorf("parsing content data: %w", err)
	}
	return data, nil
}
