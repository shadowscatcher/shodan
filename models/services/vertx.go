package services

type Vertx struct {
	// Date that the firmware was released
	FirmwareData string `json:"firmware_data"`

	// Version of the firmware
	FirmwareVersion string `json:"firmware_version"`

	// Local IP of the device
	InternalIP string `json:"internal_ip"`

	// MAC address
	MAC string `json:"mac"`

	// Name of the door controller
	Name string `json:"name"`

	// Product type
	Type string `json:"type"`
}
