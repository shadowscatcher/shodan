package models

// Result of request to "host/search" and "host/search/count" endpoints
type SearchResult struct {
	Matches []Service          `json:"matches"`
	Total   int                `json:"total"`
	Facets  map[string][]Facet `json:"facets,omitempty"`
}

type ExploitResult struct {
	Matches []Exploit          `json:"matches"`
	Total   int                `json:"total"`
	Facets  map[string][]Facet `json:"facets,omitempty"`
}

type Facet struct {
	Count int         `json:"count"`
	Value interface{} `json:"value"`
}
