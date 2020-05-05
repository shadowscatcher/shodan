package services

type SNMP struct {
	// Contact information
	Contact string `json:"contact"`

	// Description of the device
	Description string `json:"description"`

	// Physical location
	Location *string `json:"location"`

	// Name given to the device by the owner
	Name *string `json:"name"`

	Uptime       string `json:"uptime"`
	ObjectId     string `json:"objectid"`
	Services     string `json:"services"`
	OrLastChange string `json:"orlastchange"`
	OrDescr      string `json:"ordescr"`
	OrUptime     string `json:"oruptime"`
	OrId         string `json:"orid"`
}
