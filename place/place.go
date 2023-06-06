package place

import "github.com/oyvinddd/apple-maps-server-sdk/location"

type (
	Place struct {
		Country string `json:"country"`

		CountryCode string `json:"countryCode"`

		DisplayMapRegion location.MapRegion `json:"displayMapRegion"`

		FormattedAddressLines []string `json:"formattedAddressLine"`

		Name string `json:"name"`

		Coordinate location.Location `json:"coordinate"`

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
