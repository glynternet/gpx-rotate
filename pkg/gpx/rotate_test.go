package gpx_test

import (
	"testing"

	"github.com/glynternet/gpx/pkg/gpx"
	"github.com/stretchr/testify/assert"
	gpxgo "github.com/tkrajina/gpxgo/gpx"
)

func TestRotate(t *testing.T) {
	for _, tc := range []struct {
		name     string
		input    []gpxgo.GPXPoint
		n        int
		expected []gpxgo.GPXPoint
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
			assert.Equal(t, diffable(tc.expected), diffable(gpx.Rotated(tc.input, tc.n)))
		})
	}
}

func points(names ...string) []gpxgo.GPXPoint {
	ps := make([]gpxgo.GPXPoint, len(names))
	for i, name := range names {
		ps[i] = gpxgo.GPXPoint{Name: name}
	}
	return ps
}

func diffable(ps []gpxgo.GPXPoint) []string {
	vs := make([]string, len(ps))
	for i, p := range ps {
		vs[i] = p.Name
	}
	return vs
}
