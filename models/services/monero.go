package services

type Monero struct {
	Credits uint64 `json:"credits"`
	TopHash string `json:"top_hash"`

	// Number of alternative blocks to main chain
	AltBlocksCount         int    `json:"alt_blocks_count"`
	BlockSizeLimit         int    `json:"block_size_limit"`
	BlockSizeMedian        int    `json:"block_size_median,omitempty"`
	BlockWeightLimit       int    `json:"block_weight_limit,omitempty"`
	BlockWeightMedian      int    `json:"block_weight_median,omitempty"`
	BootstrapDaemonAddress string `json:"bootstrap_daemon_address,omitempty"`

	// List of nodes that are connected to the daemon
	Connections               []MoneroConnection `json:"connections"`
	CumulativeDifficulty      int                `json:"cumulative_difficulty"`
	CumulativeDifficultyTop64 int                `json:"cumulative_difficulty_top64,omitempty"`
	DatabaseSize              int                `json:"database_size,omitempty"`

	// Network difficulty
	Difficulty      int `json:"difficulty"`
	DifficultyTop64 int `json:"difficulty_top64,omitempty"`
	FreeSpace       int `json:"free_space,omitempty"`

	// Grey peerlist size
	GreyPeerlistSize int `json:"grey_peerlist_size"`

	// Length of longest chain known to daemon
	Height                 int `json:"height"`
	HeightWithoutBootstrap int `json:"height_without_bootstrap,omitempty"`

	// Number of peers connected to the node
	IncomingConnectionsCount int    `json:"incoming_connections_count"`
	Mainnet                  bool   `json:"mainnet,omitempty"`
	Nettype                  string `json:"nettype,omitempty"`
	Offline                  bool   `json:"offline,omitempty"`

	// Number of outgoing connections from the daemon
	OutgoingConnectionsCount int  `json:"outgoing_connections_count"`
	RPCConnectionsCount      int  `json:"rpc_connections_count,omitempty"`
	Stagenet                 bool `json:"stagenet,omitempty"`
	StartTime                int  `json:"start_time"`

	// RPC error code
	Status string `json:"status"`

	// Current target for next proof of work
	Target int `json:"target"`

	// Height of the next block in the chain
	TargetHeight int `json:"target_height"`

	// Whether the node is on the testnet (true) or not (false)
	Testnet bool `json:"testnet"`

	// Hash of the highest block in the chain
	TopBlockHash string `json:"top_block_hash"`

	// Total number of non-coinbase transactions in the chain
	TxCount int `json:"tx_count"`

	// Number of transactions that have been broadcast but not included in a block
	TxPoolSize           int    `json:"tx_pool_size"`
	Untrusted            bool   `json:"untrusted,omitempty"`
	UpdateAvailable      bool   `json:"update_available,omitempty"`
	Version              string `json:"version,omitempty"`
	WasBootstrapEverUsed bool   `json:"was_bootstrap_ever_used,omitempty"`

	// White peerlist size
	WhitePeerlistSize        int    `json:"white_peerlist_size"`
	WideCumulativeDifficulty string `json:"wide_cumulative_difficulty,omitempty"`
	WideDifficulty           string `json:"wide_difficulty,omitempty"`
}

type MoneroConnection struct {
	Address         string `json:"address"`
	AvgDownload     int    `json:"avg_download"`
	AvgUpload       int    `json:"avg_upload"`
	ConnectionID    string `json:"connection_id"`
	CurrentDownload int    `json:"current_download"`
	CurrentUpload   int    `json:"current_upload"`
	Height          int    `json:"height"`
	Host            string `json:"host"`
	IP              string `json:"ip"`
	Incoming        bool   `json:"incoming"`
	LiveTime        int    `json:"live_time"`
	LocalIP         bool   `json:"local_ip"`
	Localhost       bool   `json:"localhost"`
	PeerID          string `json:"peer_id"`
	Port            string `json:"port"`
	PruningSeed     int    `json:"pruning_seed,omitempty"`
	RPCPort         int    `json:"rpc_port,omitempty"`
	RecvCount       int    `json:"recv_count"`
	RecvIdleTime    int    `json:"recv_idle_time"`
	SendCount       int    `json:"send_count"`
	SendIdleTime    int    `json:"send_idle_time"`
	State           string `json:"state"`
	SupportFlags    int    `json:"support_flags"`
}
