package services

// Availability: Docker services that allow remote connections and don’t have authentication enabled
type Docker struct {
	// API version of this Docker service
	APIVersion string `json:"ApiVersion"`

	// Architecture of the server running the service
	Arch string `json:"Arch"`

	// Date and time that the build was created
	BuildTime string `json:"BuildTime,omitempty"`

	Components   []DockerComponent `json:"Components,omitempty"`
	Containers   []DockerContainer `json:"Containers,omitempty"`
	Engine       string            `json:"Engine,omitempty"`
	Experimental bool              `json:"Experimental,omitempty"`

	// Git commit identifier
	GitCommit string `json:"GitCommit"`

	// Version of Go used for the build
	GoVersion string `json:"GoVersion"`

	// Kernel version of the server
	KernelVersion string `json:"KernelVersion"`

	// Minimum API version the server supports
	MinApiversion string `json:"MinAPIVersion,omitempty"`

	EulerVersion string `json:"EulerVersion,omitempty"`

	// Host operating system
	OS         string         `json:"Os"`
	PkgVersion string         `json:"PkgVersion,omitempty"`
	Platform   DockerPlatform `json:"Platform,omitempty"`

	// Docker version
	Version string `json:"Version"`
}

type DockerComponent struct {
	Details DockerComponentDetails `json:"Details"`
	Name    string                 `json:"Name"`
	Version string                 `json:"Version"`
}

type DockerContainer struct {
	Command         string                    `json:"Command"`
	Created         int                       `json:"Created"`
	FinishedAt      int                       `json:"FinishedAt,omitempty"`
	HostConfig      DockerContainerHostConfig `json:"HostConfig,omitempty"`
	ID              string                    `json:"Id"`
	Image           string                    `json:"Image"`
	ImageID         string                    `json:"ImageID,omitempty"`
	Labels          interface{}               `json:"Labels,omitempty"`
	Mounts          interface{}               `json:"Mounts,omitempty"`
	Names           []string                  `json:"Names"`
	NetworkSettings interface{}               `json:"NetworkSettings,omitempty"`
	Ports           interface{}               `json:"Ports"`
	StartedAt       int                       `json:"StartedAt,omitempty"`
	State           string                    `json:"State,omitempty"`
	Status          string                    `json:"Status"`
}

type DockerComponentDetails struct {
	APIVersion    string `json:"ApiVersion,omitempty"`
	Arch          string `json:"Arch,omitempty"`
	BuildTime     string `json:"BuildTime,omitempty"`
	Experimental  string `json:"Experimental,omitempty"`
	GitCommit     string `json:"GitCommit"`
	GoVersion     string `json:"GoVersion,omitempty"`
	KernelVersion string `json:"KernelVersion,omitempty"`
	MinApiversion string `json:"MinAPIVersion,omitempty"`
	Os            string `json:"Os,omitempty"`
}

type DockerContainerHostConfig struct {
	NetworkMode string `json:"NetworkMode,omitempty"`
}

type DockerPlatform struct {
	Name string `json:"Name"`
}
