package services

type Cassandra struct {
	// Name of the cluster
	Name string `json:"name"`

	// List of keyspaces available
	Keyspaces []string `json:"keyspaces"`

	// Algorithm used to partition the data across the cluster
	Partitioner string `json:"partitioner"`

	// Algorithm used to determine the network topology of the cluster
	Snitch string `json:"snitch"`
}
