package services

type MQTT struct {
	// Response status code (0 for a successful response)
	Code int `json:"code"`

	// List of messages captured from the server during a 5 second window
	Messages []MqttMessage `json:"messages"`
}

type MqttMessage struct {
	// The payload is only available on system properties
	Payload *string `json:"payload"`
	Topic   string  `json:"topic"`
}
