package services

type ElasticJvmMem struct {
	DirectMaxInBytes   int `json:"direct_max_in_bytes"`
	HeapInitInBytes    int `json:"heap_init_in_bytes"`
	HeapMaxInBytes     int `json:"heap_max_in_bytes"`
	NonHeapInitInBytes int `json:"non_heap_init_in_bytes"`
	NonHeapMaxInBytes  int `json:"non_heap_max_in_bytes"`
	HeapUsedInBytes    int `json:"heap_used_in_bytes"`
}

type ElasticOsMem struct {
	FreeInBytes  int `json:"free_in_bytes,omitempty"`
	FreePercent  int `json:"free_percent,omitempty"`
	TotalInBytes int `json:"total_in_bytes"`
	UsedInBytes  int `json:"used_in_bytes,omitempty"`
	UsedPercent  int `json:"used_percent,omitempty"`
}

type ElasticProcess struct {
	ID                      int                        `json:"id"`
	MaxFileDescriptors      int                        `json:"max_file_descriptors,omitempty"`
	Mlockall                bool                       `json:"mlockall"`
	RefreshIntervalInMillis int                        `json:"refresh_interval_in_millis"`
	CPU                     ElasticCPULoad             `json:"cpu,omitempty"`
	OpenFileDescriptors     ElasticOpenFileDescriptors `json:"open_file_descriptors,omitempty"`
}
