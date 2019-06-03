package services

type Hive struct {
	Databases []HiveDatabase `json:"databases"`
}

type HiveDatabase struct {
	// name of the database
	Name string `json:"name"`

	// List of table objects
	Tables []HiveTable `json:"tables"`
}

type HiveTable struct {
	Name       string              `json:"name"`
	Properties []map[string]string `json:"properties"`
}
