package services

type InfluxDb struct {
	Uptime          string   `json:"uptime"`
	GoMaxProcs      int      `json:"go_max_procs"`
	GoVersion       string   `json:"go_version"`
	GoOS            string   `json:"go_os"`
	GoArch          string   `json:"go_arch"`
	NetworkHostname string   `json:"network_hostname"`
	Version         string   `json:"version"`
	BindAddress     string   `json:"bind_address"`
	Build           string   `json:"build"`
	Databases       []string `json:"databases"`
}
