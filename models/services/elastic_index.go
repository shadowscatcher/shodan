package services

type ElasticIndex struct {
	Primaries ElasticIndexStats `json:"primaries"`
	Total     ElasticIndexStats `json:"total"`
	UUID      string            `json:"uuid,omitempty"`
}

type ElasticIndexStats struct {
	Indexing ElasticIndexing `json:"indexing,omitempty"`
}

type ElasticIndexing struct {
	DeleteCurrent        int  `json:"delete_current"`
	DeleteTimeInMillis   int  `json:"delete_time_in_millis"`
	DeleteTotal          int  `json:"delete_total"`
	IndexCurrent         int  `json:"index_current"`
	IndexFailed          int  `json:"index_failed,omitempty"`
	IndexTimeInMillis    int  `json:"index_time_in_millis"`
	IndexTotal           int  `json:"index_total"`
	IsThrottled          bool `json:"is_throttled"`
	NoopUpdateTotal      int  `json:"noop_update_total"`
	ThrottleTimeInMillis int  `json:"throttle_time_in_millis"`
}

type ElasticCompletion struct {
	SizeInBytes int `json:"size_in_bytes"`
}

type ElasticNodeCount struct {
	Client           int `json:"client,omitempty"`
	CoordinatingOnly int `json:"coordinating_only,omitempty"`
	VotingOnly       int `json:"voting_only,omitempty"`
	Data             int `json:"data,omitempty"`
	DataOnly         int `json:"data_only,omitempty"`
	Ingest           int `json:"ingest,omitempty"`
	Master           int `json:"master,omitempty"`
	MasterData       int `json:"master_data,omitempty"`
	MasterOnly       int `json:"master_only,omitempty"`
	ML               int `json:"ml"`
	Total            int `json:"total"`
}

type ElasticCPULoad struct {
	Percent int `json:"percent"`
}

type ElasticCpuItem struct {
	CacheSizeInBytes int    `json:"cache_size_in_bytes"`
	CoresPerSocket   int    `json:"cores_per_socket"`
	Count            int    `json:"count"`
	Mhz              int    `json:"mhz"`
	Model            string `json:"model,omitempty"`
	TotalCores       int    `json:"total_cores"`
	TotalSockets     int    `json:"total_sockets"`
	Vendor           string `json:"vendor,omitempty"`
}

type ElasticIndexDocs struct {
	Count   int `json:"count"`
	Deleted int `json:"deleted"`
}

type ElasticFielddata struct {
	Evictions         int `json:"evictions"`
	MemorySizeInBytes int `json:"memory_size_in_bytes"`
}

type ElasticFilterCache struct {
	Evictions         int `json:"evictions"`
	MemorySizeInBytes int `json:"memory_size_in_bytes"`
}

type ElasticFS struct {
	AvailableInBytes     int    `json:"available_in_bytes,omitempty"`
	DiskIoOp             int    `json:"disk_io_op,omitempty"`
	DiskIoSizeInBytes    int    `json:"disk_io_size_in_bytes,omitempty"`
	DiskQueue            string `json:"disk_queue,omitempty"`
	DiskReadSizeInBytes  int    `json:"disk_read_size_in_bytes,omitempty"`
	DiskReads            int    `json:"disk_reads,omitempty"`
	DiskServiceTime      string `json:"disk_service_time,omitempty"`
	DiskWriteSizeInBytes int    `json:"disk_write_size_in_bytes,omitempty"`
	DiskWrites           int    `json:"disk_writes,omitempty"`
	FreeInBytes          int    `json:"free_in_bytes,omitempty"`
	Spins                string `json:"spins,omitempty"`
	TotalInBytes         int    `json:"total_in_bytes,omitempty"`
}

type ElasticIdcache struct {
	MemorySizeInBytes int `json:"memory_size_in_bytes"`
}

type ElasticShardIndex struct {
	Primaries   ElasticIndexMetric `json:"primaries"`
	Replication ElasticIndexMetric `json:"replication"`
	Shards      ElasticIndexMetric `json:"shards"`
}

type ElasticIndexMetric struct {
	Avg float64 `json:"avg"`
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}

type ElasticIndices struct {
	Completion  ElasticCompletion    `json:"completion"`
	Count       int                  `json:"count"`
	Docs        ElasticIndexDocs     `json:"docs"`
	Fielddata   ElasticFielddata     `json:"fielddata"`
	FilterCache ElasticFilterCache   `json:"filter_cache,omitempty"`
	IDCache     ElasticIdcache       `json:"id_cache,omitempty"`
	Percolate   ElasticPercolate     `json:"percolate,omitempty"`
	QueryCache  ElasticQueryCache    `json:"query_cache,omitempty"`
	Segments    ElasticSegments      `json:"segments"`
	Shards      ElasticIndicesShards `json:"shards"`
	Store       ElasticStore         `json:"store"`
}

type ElasticIndicesShards struct {
	Index       ElasticShardIndex `json:"index,omitempty"`
	Primaries   int               `json:"primaries,omitempty"`
	Replication float64           `json:"replication,omitempty"`
	Total       int               `json:"total,omitempty"`
}

type ElasticJVMdata struct {
	MaxUptimeInMillis int                 `json:"max_uptime_in_millis"`
	Mem               ElasticJvmMem       `json:"mem"`
	Threads           int                 `json:"threads"`
	Versions          []ElasticJvmVersion `json:"versions,omitempty"`
}

type ElasticJvmVersion struct {
	Count     int    `json:"count"`
	VMName    string `json:"vm_name"`
	VMVendor  string `json:"vm_vendor"`
	VMVersion string `json:"vm_version"`
	Version   string `json:"version"`
}

type ElasticOSname struct {
	Count int    `json:"count"`
	Name  string `json:"name,omitempty"`
}

type ElasticNetworkTypes struct {
	HTTPTypes      map[string]interface{} `json:"http_types"`
	TransportTypes map[string]interface{} `json:"transport_types"`
}

type ElasticNodes struct {
	Count          ElasticNodeCount       `json:"count"`
	FS             ElasticFS              `json:"fs"`
	JVM            ElasticJVMdata         `json:"jvm"`
	NetworkTypes   ElasticNetworkTypes    `json:"network_types,omitempty"`
	OS             ElasticOS              `json:"os"`
	Plugins        []ElasticPlugin        `json:"plugins,omitempty"`
	Process        ElasticProcess         `json:"process"`
	Versions       []string               `json:"versions"`
	Ingest         ElasticIngest          `json:"ingest,omitempty"`
	PackagingTypes []PackagingType        `json:"packaging_types"`
	DiscoveryTypes map[string]interface{} `json:"discovery_types"`
}

type ElasticOpenFileDescriptors struct {
	Avg int `json:"avg"`
	Max int `json:"max"`
	Min int `json:"min"`
}

type ElasticOS struct {
	AllocatedProcessors int                 `json:"allocated_processors,omitempty"`
	AvailableProcessors int                 `json:"available_processors"`
	CPU                 []ElasticCpuItem    `json:"cpu,omitempty"`
	Mem                 ElasticOsMem        `json:"mem"`
	Names               []ElasticOSname     `json:"names,omitempty"`
	PrettyNames         []ElasticPrettyName `json:"pretty_names,omitempty"`
}

type ElasticPercolate struct {
	Current           int    `json:"current"`
	MemorySize        string `json:"memory_size"`
	MemorySizeInBytes int    `json:"memory_size_in_bytes"`
	Queries           int    `json:"queries"`
	TimeInMillis      int    `json:"time_in_millis"`
	Total             int    `json:"total"`
}

type ElasticPrettyName struct {
	Count      int    `json:"count"`
	PrettyName string `json:"pretty_name"`
}

type ElasticQueryCache struct {
	CacheCount        int `json:"cache_count"`
	CacheSize         int `json:"cache_size"`
	Evictions         int `json:"evictions"`
	HitCount          int `json:"hit_count"`
	MemorySizeInBytes int `json:"memory_size_in_bytes"`
	MissCount         int `json:"miss_count"`
	TotalCount        int `json:"total_count"`
}

type ElasticSegments struct {
	Count                       int                    `json:"count"`
	DocValuesMemoryInBytes      int                    `json:"doc_values_memory_in_bytes,omitempty"`
	FileSizes                   map[string]interface{} `json:"file_sizes,omitempty"`
	FixedBitSetMemoryInBytes    int                    `json:"fixed_bit_set_memory_in_bytes"`
	IndexWriterMaxMemoryInBytes int                    `json:"index_writer_max_memory_in_bytes,omitempty"`
	IndexWriterMemoryInBytes    int                    `json:"index_writer_memory_in_bytes"`
	MaxUnsafeAutoIDTimestamp    int                    `json:"max_unsafe_auto_id_timestamp,omitempty"`
	MemoryInBytes               int                    `json:"memory_in_bytes"`
	NormsMemoryInBytes          int                    `json:"norms_memory_in_bytes,omitempty"`
	PointsMemoryInBytes         int                    `json:"points_memory_in_bytes,omitempty"`
	StoredFieldsMemoryInBytes   int                    `json:"stored_fields_memory_in_bytes,omitempty"`
	TermVectorsMemoryInBytes    int                    `json:"term_vectors_memory_in_bytes,omitempty"`
	TermsMemoryInBytes          int                    `json:"terms_memory_in_bytes,omitempty"`
	TermsOffheapMemoryInBytes   int                    `json:"terms_offheap_memory_in_bytes,omitempty"`
	VersionMapMemoryInBytes     int                    `json:"version_map_memory_in_bytes"`
}

type ElasticStore struct {
	SizeInBytes          int `json:"size_in_bytes"`
	ThrottleTimeInMillis int `json:"throttle_time_in_millis,omitempty"`
}

type PackagingType struct {
	Count  int    `json:"count"`
	Flavor string `json:"flavor"`
	Type   string `json:"type"`
}
