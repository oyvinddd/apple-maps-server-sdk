package location

import "fmt"

type (
	Location struct {
		Latitude float32 `json:"latitude"`

		Longitude float32 `json:"longitude"`
	}

	Region struct {
		UpperRight Location

		LowerLeft Location
	}

	MapRegion struct {
		// EastLongitude A double value that describes the east longitude of the map region.
		EastLongitude float32 `json:"eastLongitude"`

		// NorthLatitude A double value that describes the north latitude of the map region.
		NorthLatitude float32 `json:"northLatitude"`

		// SouthLatitude A double value that describes the south latitude of the map region.
		SouthLatitude float32 `json:"southLatitude"`

		// WestLongitude A double value that describes west longitude of the map region.
		WestLongitude float32 `json:"westLongitude"`
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
