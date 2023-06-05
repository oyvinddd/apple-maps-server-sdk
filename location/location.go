package location

type Location struct {
	Latitude float32

	Longitude float32
}

// New creates a new coordinate object
func New(lat, long float32) Location {
	return Location{
		Latitude:  lat,
		Longitude: long,
	}
}
