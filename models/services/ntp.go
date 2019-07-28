package services

import (
	"encoding/json"
)

var documentedKeys = []string{
	"monlist", "system", "version", "clock", "clock_offset", "delay", "frequency", "jitter", "leap", "noise", "offset",
	"poll", "precision", "reftime", "root_delay", "rootdelay", "rootdisp", "stability", "stratum",
}

type NTP struct {
	// Results from sending a “monlist” command, which returns a list of recently-seen clients
	Monlist interface{} `json:"monlist"`
	// Operating system
	System string `json:"system,omitempty"`
	// Full version information
	Version interface{} `json:"version"`

	RefId       string      `json:"refid"`
	State       int         `json:"state"`
	Clock       string      `json:"clock,omitempty"`
	ClockOffset float64     `json:"clock_offset"`
	Delay       float64     `json:"delay"`
	MinTC       int         `json:"mintc,omitempty"`
	TC          int         `json:"tc,omitempty"`
	Peer        uint64      `json:"peer,omitempty"`
	Processor   string      `json:"processor,omitempty"`
	Frequency   interface{} `json:"frequency,omitempty"`
	Jitter      float64     `json:"jitter,omitempty"`
	SysJitter   interface{} `json:"sys_jitter,omitempty"`
	ClkJitter   interface{} `json:"clk_jitter,omitempty"`
	ClkWander   interface{} `json:"clk_wander,omitempty"`
	Phase       interface{} `json:"phase"`
	Leap        int         `json:"leap"`
	Noise       float64     `json:"noise,omitempty"`
	Offset      interface{} `json:"offset,omitempty"`
	Poll        int         `json:"poll"`
	Precision   int         `json:"precision"`
	Reftime     interface{} `json:"reftime"`
	RootDelay   float64     `json:"root_delay"`
	Rootdelay   interface{} `json:"rootdelay,omitempty"`
	RootDisp    interface{} `json:"rootdisp,omitempty"`
	Stability   float64     `json:"stability,omitempty"`
	Stratum     int         `json:"stratum"`
	Extra       map[string]interface{}
}

type ntpOverhead NTP

func (n *NTP) UnmarshalJSON(bytes []byte) (err error) {
	overhead := ntpOverhead{}

	if err = json.Unmarshal(bytes, &overhead); err == nil {
		*n = NTP(overhead)
	}

	extraValues := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &extraValues); err == nil {
		for _, key := range documentedKeys {
			delete(extraValues, key)
		}
		n.Extra = extraValues
	}

	return err
}

// todo: loss of Extra data
//func (n *NTP) MarshalJSON() ([]byte, error) {
//	panic("Not implemented")
//}
