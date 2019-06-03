package services

type Elastic struct {
	// General information about the cluster
	Cluster ElasticCluster `json:"cluster"`

	// List of nodes/peers
	Nodes ElasticNode `json:"nodes"`

	// List of available indexes
	Indices map[string]ElasticIndex `json:"indices"`
}

type ElasticCluster struct {
	ClusterName string           `json:"cluster_name,omitempty"`
	ClusterUUID string           `json:"cluster_uuid,omitempty"`
	Indices     ElasticIndices   `json:"indices,omitempty"`
	Nodes       ElasticNodes     `json:"nodes,omitempty"`
	NodesMore   ElasticNodesMore `json:"_nodes,omitempty"`
	Status      string           `json:"status,omitempty"`
	Timestamp   int              `json:"timestamp,omitempty"`
}

type ElasticNodesMore struct {
	Failed     int              `json:"failed"`
	Failures   []ElasticFailure `json:"failures,omitempty"`
	Successful int              `json:"successful"`
	Total      int              `json:"total"`
}

type ElasticFailure struct {
	CausedBy ElasticFailureCausedBy `json:"caused_by"`
	NodeID   string                 `json:"node_id"`
	Reason   string                 `json:"reason"`
	Type     string                 `json:"type"`
}

type ElasticFailureCausedBy struct {
	CausedBy *ElasticFailureCausedBy `json:"caused_by,omitempty"`
	Reason   string                  `json:"reason"`
	Type     string                  `json:"type"`
}
