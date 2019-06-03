package models

// An object containing all of the location information for the device.
type Location struct {
	// The latitude for the geolocation of the device
	Latitude *float32 `json:"latitude"`

	// The longitude for the geolocation of the device
	Longitude *float32 `json:"longitude"`

	//  The name of the city where the device is located
	City *string `json:"city"`

	// The 2-letter country code for the device location
	CountryCode *string `json:"country_code"`

	// The 3-letter country code for the device location
	CountryCode3 *string `json:"country_code3"`

	// The name of the country where the device is located
	CountryName *string `json:"country_name"`

	// The area code for the device's location. Only available for the US
	AreaCode *int `json:"area_code"`

	// The name of the region where the device is located
	RegionCode *string `json:"region_code"`

	// The designated market area code for the area where the device is located. Only available for the US
	DmaCode *int `json:"dma_code"`

	// The postal code for the device's location
	PostalCode *string `json:"postal_code"`
}
