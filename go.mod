module github.com/glynternet/gpx

go 1.17

require (
	github.com/glynternet/pkg v0.0.2
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.7.0
	github.com/tkrajina/gpxgo v1.1.2
)

require github.com/go-kit/kit v0.12.0 // indirect

replace github.com/tkrajina/gpxgo => github.com/glynternet/gpxgo v0.0.0-20221016115515-314c84573cbc
