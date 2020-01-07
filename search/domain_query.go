package search

type DomainQuery struct {
	Domain     string
	History    bool
	RecordType string
	Page       int // pages after first require enterprise access
}
