package place

import "github.com/oyvinddd/apple-maps-server-sdk/location"

type Place struct {
	Country string

	CountryCode string

	DisplayMapRegion location.Region

	FormattedAddressLines []string

	Name string

	Coordinate location.Location

	StructuredAddress string // TODO: ....
}
