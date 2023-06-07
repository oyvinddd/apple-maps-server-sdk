package location

import "fmt"

type (
	Location struct {
		Latitude float32 `json:"latitude"`

		Longitude float32 `json:"longitude"`
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

	Place struct {
		Country string `json:"country"`

		CountryCode string `json:"countryCode"`

		DisplayMapRegion MapRegion `json:"displayMapRegion"`

		FormattedAddressLines []string `json:"formattedAddressLine"`

		Name string `json:"name"`

		Coordinate Location `json:"coordinate"`

		StructuredAddress StructuredAddress `json:"structuredAddress"`
	}

	StructuredAddress struct {
		AdministrativeArea string `json:"administrativeArea"`

		AdministrativeAreaCode string `json:"administrativeAreaCode"`

		AreasOfInterest []string `json:"areasOfInterest"`

		DependentLocalities []string `json:"dependentLocalities"`

		FullThoroughfare string `json:"fullThoroughfare"`

		Locality string `json:"locality"`

		PostCode string `json:"postCode"`

		SubLocality string `json:"subLocality"`

		// SubThoroughfare The short code for the state or area.
		SubThoroughfare string `json:"subThoroughfare"`

		// Thoroughfare The state or province of the place
		Thoroughfare string `json:"thoroughfare"`
	}
)

// New creates a new coordinate object
func New(lat, lng float32) Location {
	return Location{Latitude: lat, Longitude: lng}
}

func (l Location) String() string {
	return fmt.Sprintf("%f,%f", l.Latitude, l.Longitude)
}
