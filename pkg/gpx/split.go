package gpx

import (
	gpxgo "github.com/tkrajina/gpxgo/gpx"
)

func Split(gpx gpxgo.GPX) []gpxgo.GPX {
	tracksCount := len(gpx.Tracks)
	if tracksCount == 0 || tracksCount == 1 {
		return []gpxgo.GPX{gpx}
	}
	var split []gpxgo.GPX
	for _, track := range gpx.Tracks {
		c := gpx
		c.Name = track.Name
		c.Tracks = []gpxgo.GPXTrack{track}
		split = append(split, c)
	}
	return split
}
