package services

type Mongo struct {
	// Whether or not the database has any authentication enabled. Note that the server may only have partial
	// authentication and allow certain commands to work without credentials
	Authentication bool `json:"authentication"`

	// The result of running the “buildInfo” command. Summary information about the current build
	BuildInfo MongoBuildInfo `json:"buildInfo"`

	// The result of running the “listDatabases” command
	ListDatabases MongoListDatabases `json:"listDatabases,omitempty"`

	// The result of running the “serverStatus” command. An overview of the database’s state
	ServerStatus map[string]interface{} `json:"serverStatus,omitempty"`
}

type MongoListDatabases struct {
	Databases             []MongoDatabase `json:"databases"`
	Ok                    float64         `json:"ok"`
	TotalSize             float64         `json:"totalSize"`
	TotalUncompressedSize float64         `json:"totalUncompressedSize,omitempty"`
}

type MongoBuildInfo struct {
	Allocator         string                `json:"allocator,omitempty"`
	Bits              int                   `json:"bits"`
	BuildEnvironment  MongoBuildEnvironment `json:"buildEnvironment,omitempty"`
	CompilerFlags     string                `json:"compilerFlags,omitempty"`
	CompilerName      string                `json:"compiler name,omitempty"`
	CompilerVersion   string                `json:"compiler version,omitempty"`
	Debug             bool                  `json:"debug,omitempty"`
	GitVersion        string                `json:"gitVersion"`
	JavascriptEngine  string                `json:"javascriptEngine,omitempty"`
	LoaderFlags       string                `json:"loaderFlags,omitempty"`
	MaxBsonObjectSize int                   `json:"maxBsonObjectSize,omitempty"`
	MemorySanitize    bool                  `json:"memory_sanitize,omitempty"`
	Modules           []string              `json:"modules,omitempty"`
	Ok                float64               `json:"ok"`
	OpenSslversion    string                `json:"OpenSSLVersion,omitempty"`
	Openssl           MongoOpenSSl          `json:"openssl,omitempty"`
	PcreJit           bool                  `json:"pcre-jit,omitempty"`
	PsmdbVersion      string                `json:"psmdbVersion,omitempty"`
	SonarVersion      string                `json:"sonarVersion,omitempty"`
	Sonardb           bool                  `json:"sonardb,omitempty"`
	StorageEngines    []string              `json:"storageEngines,omitempty"`
	SysInfo           string                `json:"sysInfo"`
	TargetMinOs       string                `json:"targetMinOS,omitempty"`
	Timestamp         string                `json:"timestamp,omitempty"`
	TokukvVersion     string                `json:"tokukvVersion,omitempty"`
	TokumxVersion     string                `json:"tokumxVersion,omitempty"`
	Version           string                `json:"version"`
	VersionArray      []int                 `json:"versionArray,omitempty"`
}

type MongoBuildEnvironment struct {
	Bits       int    `json:"bits,omitempty"`
	Cc         string `json:"cc,omitempty"`
	Ccflags    string `json:"ccflags,omitempty"`
	Cxx        string `json:"cxx,omitempty"`
	Cxxflags   string `json:"cxxflags,omitempty"`
	Distarch   string `json:"distarch,omitempty"`
	Distmod    string `json:"distmod,omitempty"`
	Linkflags  string `json:"linkflags,omitempty"`
	TargetArch string `json:"target_arch,omitempty"`
	TargetOs   string `json:"target_os"`
}

type MongoDatabase struct {
	Collections []string `json:"collections"`
	Empty       bool     `json:"empty,omitempty"`
	Name        string   `json:"name"`
	Size        float64  `json:"size,omitempty"`
	SizeOnDisk  float64  `json:"sizeOnDisk"`
}

type MongoOpenSSl struct {
	Compiled string `json:"compiled,omitempty"`
	Running  string `json:"running"`
}
