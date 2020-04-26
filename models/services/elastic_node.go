package services

type ElasticNode struct {
	ClusterName string                     `json:"cluster_name"`
	Nodes       map[string]ElasticNodeInfo `json:"nodes"`
	NodesStat   ElasticNodeStat            `json:"_nodes,omitempty"`
}

type ElasticNodeStat struct {
	Failed     int              `json:"failed"`
	Failures   []ElasticFailure `json:"failures,omitempty"`
	Successful int              `json:"successful"`
	Total      int              `json:"total"`
}

type ElasticAttributes struct {
	AwsAvailabilityZone  string `json:"aws_availability_zone,omitempty"`
	BoxType              string `json:"box_type,omitempty"`
	Client               string `json:"client,omitempty"`
	Data                 string `json:"data,omitempty"`
	FaultDomain          string `json:"fault_domain,omitempty"`
	Local                string `json:"local,omitempty"`
	Master               string `json:"master,omitempty"`
	MaxLocalStorageNodes string `json:"max_local_storage_nodes,omitempty"`
	MlEnabled            string `json:"ml.enabled,omitempty"`
	MlMachineMemory      string `json:"ml.machine_memory,omitempty"`
	MlMaxOpenJobs        string `json:"ml.max_open_jobs,omitempty"`
	Rack                 string `json:"rack,omitempty"`
	Role                 string `json:"role,omitempty"`
	UpdateDomain         string `json:"update_domain,omitempty"`
	XpackInstalled       string `json:"xpack.installed,omitempty"`
}

type ElasticCPUdata struct {
	CacheSizeInBytes int    `json:"cache_size_in_bytes,omitempty"`
	CoresPerSocket   int    `json:"cores_per_socket"`
	Mhz              int    `json:"mhz,omitempty"`
	Model            string `json:"model,omitempty"`
	TotalCores       int    `json:"total_cores"`
	TotalSockets     int    `json:"total_sockets"`
	Vendor           string `json:"vendor,omitempty"`
}

type ElasticNodeHTTP struct {
	BoundAddress            interface{} `json:"bound_address"`
	MaxContentLengthInBytes int         `json:"max_content_length_in_bytes"`
	PublishAddress          string      `json:"publish_address"`
}

type ElasticIngest struct {
	Processors        []ElasticProcessor     `json:"processors"`
	ProcessorStats    map[string]interface{} `json:"processor_stats"`
	NumberOfPipelines int                    `json:"number_of_pipelines"`
}

type ElasticJVM struct {
	GcCollectors                          []string      `json:"gc_collectors,omitempty"`
	InputArguments                        []string      `json:"input_arguments,omitempty"`
	Mem                                   ElasticJvmMem `json:"mem,omitempty"`
	MemoryPools                           []string      `json:"memory_pools,omitempty"`
	Pid                                   int           `json:"pid,omitempty"`
	StartTimeInMillis                     int           `json:"start_time_in_millis,omitempty"`
	UsingCompressedOrdinaryObjectPointers string        `json:"using_compressed_ordinary_object_pointers,omitempty"`
	VMName                                string        `json:"vm_name,omitempty"`
	VMVendor                              string        `json:"vm_vendor,omitempty"`
	VMVersion                             string        `json:"vm_version,omitempty"`
	Version                               string        `json:"version,omitempty"`
}

type ElasticModule struct {
	Classname            string   `json:"classname"`
	Description          string   `json:"description"`
	ElasticsearchVersion string   `json:"elasticsearch_version,omitempty"`
	ExtendedPlugins      []string `json:"extended_plugins,omitempty"`
	HasNativeController  bool     `json:"has_native_controller,omitempty"`
	Isolated             bool     `json:"isolated,omitempty"`
	JavaVersion          string   `json:"java_version,omitempty"`
	Jvm                  bool     `json:"jvm,omitempty"`
	Name                 string   `json:"name"`
	RequiresKeystore     bool     `json:"requires_keystore,omitempty"`
	Site                 bool     `json:"site,omitempty"`
	Version              string   `json:"version"`
}

type ElasticNetwork struct {
	PrimaryInterface        ElasticPrimaryInterface `json:"primary_interface,omitempty"`
	RefreshIntervalInMillis int                     `json:"refresh_interval_in_millis"`
}

type ElasticNodeInfo struct {
	ThreadPool          map[string]interface{} `json:"thread_pool,omitempty"`
	Settings            map[string]interface{} `json:"settings,omitempty"`
	Attributes          ElasticAttributes      `json:"attributes,omitempty"`
	Build               string                 `json:"build,omitempty"`
	BuildFlavor         string                 `json:"build_flavor,omitempty"`
	BuildHash           string                 `json:"build_hash,omitempty"`
	BuildType           string                 `json:"build_type,omitempty"`
	HTTP                ElasticNodeHTTP        `json:"http,omitempty"`
	HTTPAddress         string                 `json:"http_address,omitempty"`
	Host                string                 `json:"host,omitempty"`
	IP                  string                 `json:"ip,omitempty"`
	Ingest              ElasticIngest          `json:"ingest,omitempty"`
	JVM                 ElasticJVM             `json:"jvm"`
	Modules             []ElasticModule        `json:"modules,omitempty"`
	Name                string                 `json:"name"`
	Network             ElasticNetwork         `json:"network,omitempty"`
	OS                  ElasticOsInfo          `json:"os"`
	Plugins             []ElasticPlugin        `json:"plugins,omitempty"`
	Process             ElasticProcess         `json:"process"`
	Roles               []string               `json:"roles,omitempty"`
	TotalIndexingBuffer int                    `json:"total_indexing_buffer,omitempty"`
	Transport           ElasticTransport       `json:"transport,omitempty"`
	TransportAddress    string                 `json:"transport_address,omitempty"`
	Version             string                 `json:"version"`
}

type ElasticOsInfo struct {
	AllocatedProcessors     int            `json:"allocated_processors,omitempty"`
	Arch                    string         `json:"arch,omitempty"`
	AvailableProcessors     int            `json:"available_processors"`
	CPU                     ElasticCPUdata `json:"cpu,omitempty"`
	Mem                     ElasticOsMem   `json:"mem,omitempty"`
	Name                    string         `json:"name,omitempty"`
	PrettyName              string         `json:"pretty_name,omitempty"`
	RefreshIntervalInMillis int            `json:"refresh_interval_in_millis"`
	Swap                    ElasticSwap    `json:"swap,omitempty"`
	Version                 string         `json:"version,omitempty"`
}

type ElasticPlugin struct {
	Classname            string   `json:"classname,omitempty"`
	Description          string   `json:"description"`
	ElasticsearchVersion string   `json:"elasticsearch_version,omitempty"`
	ExtendedPlugins      []string `json:"extended_plugins,omitempty"`
	HasNativeController  bool     `json:"has_native_controller,omitempty"`
	Isolated             bool     `json:"isolated,omitempty"`
	JavaVersion          string   `json:"java_version,omitempty"`
	JVM                  bool     `json:"jvm,omitempty"`
	Name                 string   `json:"name"`
	RequiresKeystore     bool     `json:"requires_keystore,omitempty"`
	Site                 bool     `json:"site,omitempty"`
	URL                  string   `json:"url,omitempty"`
	Version              string   `json:"version"`
}

type ElasticPrimaryInterface struct {
	Address    string `json:"address"`
	MacAddress string `json:"mac_address"`
	Name       string `json:"name"`
}

type ElasticProcessor struct {
	Type string `json:"type"`
}

type ElasticSwap struct {
	TotalInBytes int `json:"total_in_bytes"`
}

type ElasticTransport struct {
	BoundAddress   interface{}            `json:"bound_address"`
	Profiles       map[string]interface{} `json:"profiles,omitempty"`
	PublishAddress string                 `json:"publish_address"`
}
