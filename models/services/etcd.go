package services

type Etcd struct {
	// List of connected clients
	ClientUrls []string `json:"clientURLs"`

	// Unique ID of the node
	ID string `json:"id"`

	// General information about the leader of the cluster
	LeaderInfo EtcdLeaderInfo `json:"leaderInfo"`

	// Name of the cluster
	Name string `json:"name"`

	// List of connected peers
	PeerUrls []string `json:"peerURLs"`

	// Number of received append requests
	RecvAppendRequestCnt int     `json:"recvAppendRequestCnt"`
	RecvBandwidthRate    float64 `json:"recvBandwidthRate,omitempty"`
	RecvPkgRate          float64 `json:"recvPkgRate,omitempty"`

	// Number of sent append requests
	SendAppendRequestCnt int     `json:"sendAppendRequestCnt"`
	SendBandwidthRate    float64 `json:"sendBandwidthRate,omitempty"`
	SendPkgRate          float64 `json:"sendPkgRate,omitempty"`

	// Start time of the service
	StartTime string `json:"startTime"`

	// State of the node that Shodan connected to
	State   string `json:"state"`
	Version string `json:"version"`
}

type EtcdLeaderInfo struct {
	Leader    string `json:"leader"`
	StartTime string `json:"startTime"`
	Uptime    string `json:"uptime"`
}
