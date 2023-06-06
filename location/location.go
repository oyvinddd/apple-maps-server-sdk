package location

import "fmt"

type (
	Location struct {
		Latitude float32

		Longitude float32
	}

	Region struct {
		UpperRight Location

		LowerLeft Location
	}
)

// New creates a new coordinate object
func New(lat, lng float32) Location {
	return Location{Latitude: lat, Longitude: lng}
}

func (l Location) String() string {
	return fmt.Sprintf("%f,%f", l.Latitude, l.Longitude)
}

func NewRegion(upperRight, lowerLeft Location) Region {
	return Region{upperRight, lowerLeft}
}
