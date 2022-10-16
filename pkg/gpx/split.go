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
		// remove all extensions, they sometimes cause import issues and have never been useful
		track.Extensions = gpxgo.Extension{}
		for i := range track.Segments {
			// remove all extensions, they sometimes cause import issues and have never been useful
			track.Segments[i].Extensions = gpxgo.Extension{}
			for j := range track.Segments[i].Points {
				// remove all extensions, they sometimes cause import issues and have never been useful
				track.Segments[i].Points[j].Extensions = gpxgo.Extension{}
			}
		}
		c.Tracks = []gpxgo.GPXTrack{track}
		split = append(split, c)
	}
	return split
}
