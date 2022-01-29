package gpx

import "github.com/tkrajina/gpxgo/gpx"

func Rotated(ps []gpx.GPXPoint, rotation int) []gpx.GPXPoint {
	pCount := len(ps)
	if pCount == 0 {
		return ps
	}

	rotationMod := rotation % pCount
	return append(ps[(pCount-rotationMod)%pCount:], ps[:(pCount-rotationMod)%pCount]...)
}
