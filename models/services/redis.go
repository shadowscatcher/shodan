package services

type Redis struct {
	// CPU usage information
	CPU RedisCpuData `json:"cpu"`

	// List of clients that are currently connected. Note: may be one object instead of object list
	Clients interface{} `json:"clients"`

	// Cluster configuration settings – not always available on Redis instances that aren’t configured
	// to run in a cluster
	Cluster interface{} `json:"cluster,omitempty"`
	Keys    RedisKeys   `json:"keys,omitempty"`

	// Information about the keyspaces on the server
	Keyspaces map[string]string `json:"keyspace"`

	// Memory usage information
	Memory    map[string]interface{} `json:"memory"`
	Pacluster map[string]interface{} `json:"pacluster,omitempty"`

	// Persistence settings for the service
	Persistence map[string]interface{} `json:"persistence,omitempty"`

	// Replication usage and settings information
	Replication map[string]interface{} `json:"replication,omitempty"`

	// General information about the service
	Server        RedisServer        `json:"server"`
	SSL           *RedisSSL          `json:"ssl,omitempty"`
	OomPrevention RedisOomPrevention `json:"oom-prevention"`

	// Miscellaneous statistics and usage information
	Stats map[string]interface{} `json:"stats,omitempty"`
}

type RedisCpuData struct {
	UsedCPUSys          float64 `json:"used_cpu_sys"`
	UsedCPUSysChildren  float64 `json:"used_cpu_sys_children"`
	UsedCPUUser         float64 `json:"used_cpu_user"`
	UsedCPUUserChildren float64 `json:"used_cpu_user_children"`
}

type RedisKeys struct {
	// list of keys available on Redis
	Data []string `json:"data"`

	//  true if more keys are available but weren’t returned on the initial request
	More bool `json:"more"`
}

type RedisServer struct {
	ArchBits        int         `json:"arch_bits"`
	AtomicvarAPI    string      `json:"atomicvar_api,omitempty"`
	ConfigFile      string      `json:"config_file,omitempty"`
	ConfiguredHz    int         `json:"configured_hz,omitempty"`
	Executable      string      `json:"executable,omitempty"`
	GccVersion      string      `json:"gcc_version,omitempty"`
	Hz              int         `json:"hz,omitempty"`
	LruClock        int         `json:"lru_clock"`
	MultiplexingAPI string      `json:"multiplexing_api"`
	Os              string      `json:"os"`
	ProcessID       int         `json:"process_id"`
	RedisBuildID    interface{} `json:"redis_build_id,omitempty"`
	RedisGitDirty   int         `json:"redis_git_dirty"`
	RedisGitSHA1    interface{} `json:"redis_git_sha1"`
	RedisMode       string      `json:"redis_mode"`
	RedisVersion    string      `json:"redis_version"`
	RlecVersion     string      `json:"rlec_version,omitempty"`
	RunID           string      `json:"run_id"`
	TCPPort         int         `json:"tcp_port"`
	UptimeInDays    int         `json:"uptime_in_days"`
	UptimeInSeconds int         `json:"uptime_in_seconds"`
}

type RedisSSL struct {
	SSLConnectionsToCurrentCertificate  int    `json:"ssl_connections_to_current_certificate"`
	SSLConnectionsToPreviousCertificate int    `json:"ssl_connections_to_previous_certificate"`
	SSLCurrentCertificateNotAfterDate   string `json:"ssl_current_certificate_not_after_date"`
	SSLCurrentCertificateNotBeforeDate  string `json:"ssl_current_certificate_not_before_date"`
	SSLCurrentCertificateSerial         int    `json:"ssl_current_certificate_serial"`
	SSLEnabled                          string `json:"ssl_enabled"`
}

type RedisOomPrevention struct {
	On                                 string `json:"oom_prevention_on"`
	PeakUsedMemoryTotal                uint64 `json:"peak_used_memory_total"`
	PreventionThreshold                uint64 `json:"oom_prevention_threshold"`
	UsedMemoryRdb                      uint64 `json:"used_memory_rdb"`
	UsedMemoryAof                      uint64 `json:"used_memory_aof"`
	UsedMemoryTotal                    uint64 `json:"used_memory_total"`
	CurrentUsecondsWithOomPreventionOn uint64 `json:"current_useconds_with_oom_prevention_on"`
	TotalUsecondsWithOomPreventionOn   uint64 `json:"total_useconds_with_oom_prevention_on"`
	ThresholdHuman                     string `json:"oom_prevention_threshold_human"`
	UsedMemoryRdbHuman                 string `json:"used_memory_rdb_human"`
	UsedMemoryAofHuman                 string `json:"used_memory_aof_human"`
	UsedMemoryTotalHuman               string `json:"used_memory_total_human"`
	PeakUsedMemoryTotalHuman           string `json:"peak_used_memory_total_human"`
}
