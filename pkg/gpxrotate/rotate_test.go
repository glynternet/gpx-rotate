package gpxrotate_test

import (
	"testing"

	"github.com/glynternet/gpx-rotate/pkg/gpxrotate"
	"github.com/stretchr/testify/assert"
	"github.com/tkrajina/gpxgo/gpx"
)

func TestRotate(t *testing.T) {
	for _, tc := range []struct {
		name     string
		input    []gpx.GPXPoint
		n        int
		expected []gpx.GPXPoint
	}{{
		name: "no input, yields no output",
	}, {
		name:     "multiple elements, no rotation, yields same output",
		input:    points("a", "b", "c"),
		expected: points("a", "b", "c"),
	}, {
		name:     "single rotation",
		n:        1,
		input:    points("a", "b", "c"),
		expected: points("c", "a", "b"),
	}, {
		name:     "multiple rotation",
		n:        2,
		input:    points("a", "b", "c"),
		expected: points("b", "c", "a"),
	}, {
		name:     "rotate by more than slice length",
		n:        4,
		input:    points("a", "b", "c"),
		expected: points("c", "a", "b"),
	}, {
		name:     "single negative rotation",
		n:        -1,
		input:    points("a", "b", "c"),
		expected: points("b", "c", "a"),
	}, {
		name:     "negative rotation more than slice length",
		n:        -4,
		input:    points("a", "b", "c"),
		expected: points("b", "c", "a"),
	}} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, diffable(tc.expected), diffable(gpxrotate.Rotated(tc.input, tc.n)))
		})
	}
}

func points(names ...string) []gpx.GPXPoint {
	ps := make([]gpx.GPXPoint, len(names))
	for i, name := range names {
		ps[i] = gpx.GPXPoint{Name: name}
	}
	return ps
}

func diffable(ps []gpx.GPXPoint) []string {
	vs := make([]string, len(ps))
	for i, p := range ps {
		vs[i] = p.Name
	}
	return vs
}
