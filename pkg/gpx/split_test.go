package gpx_test

import (
	"testing"

	"github.com/glynternet/gpx/pkg/gpx"
	"github.com/stretchr/testify/assert"
	gpxgo "github.com/tkrajina/gpxgo/gpx"
)

func TestSplit(t *testing.T) {
	t.Run("zero tracks returns slice of same file", func(t *testing.T) {
		in := gpxgo.GPX{
			Name:   "foo",
			Tracks: nil,
		}
		out := gpx.Split(in)
		assert.Equal(t, []gpxgo.GPX{in}, out)
	})

	t.Run("single track returns slice of same file", func(t *testing.T) {
		in := gpxgo.GPX{
			Name:   "foo",
			Tracks: []gpxgo.GPXTrack{{}},
		}
		out := gpx.Split(in)
		assert.Equal(t, []gpxgo.GPX{in}, out)
	})

	t.Run("multiple tracks returns slice of individual tracks parent GPX named to match track", func(t *testing.T) {
		in := gpxgo.GPX{
			Name:   "foo",
			Tracks: []gpxgo.GPXTrack{{Name: "bar"}, {Name: "baz"}},
		}
		out := gpx.Split(in)
		assert.Equal(t, []gpxgo.GPX{
			{
				Name:   "bar",
				Tracks: []gpxgo.GPXTrack{{Name: "bar"}},
			}, {
				Name:   "baz",
				Tracks: []gpxgo.GPXTrack{{Name: "baz"}},
			},
		}, out)
	})
}
