package search

import (
	"fmt"
	"net/url"
	"strings"
)

type Params struct {
	// Structured search query
	Query Query

	// A comma-separated list of properties to get summary information on. Property names can also be in the format of
	// "property:count", where "count" is the number of facets that will be returned for a property
	// (i.e. "country:100" to get the top 100 countries for a search query)
	Facets []string

	// Whether or not to truncate some of the larger fields (default: true)
	Minify bool

	// The page number to page through results 100 at a time (default: 1)
	Page   uint
	Offset uint
}

func (p *Params) String() string {
	return fmt.Sprintf("%s; Parameters: Page: %d Offset: %d Minify: %t Facets: [%s]",
		p.Query.String(), p.Page, p.Offset, p.Minify, strings.Join(p.Facets, ","),
	)
}

func (p *Params) ToURLValues() url.Values {
	values := make(url.Values)
	values.Add("query", p.Query.String())

	if len(p.Facets) > 0 {
		values.Add("facets", strings.Join(p.Facets, ","))
	}

	if p.Minify {
		values.Add("minify", "true")
	}

	if p.Page > 1 {
		values.Add("page", fmt.Sprint(p.Page))
	}

	if p.Offset > 0 {
		values.Add("offset", fmt.Sprint(p.Offset))
	}

	return values
}

type HostParams struct {
	IP      string
	Minify  bool
	History bool
}

func (h *HostParams) ToURLValues() url.Values {
	result := make(url.Values)
	if h.Minify {
		result.Set("minify", "true")
	}

	if h.History {
		result.Set("history", "true")
	}
	return result
}

type ExploitParams struct {
	Query  ExploitsQuery
	Facets []string
	Page   int
}

func (p *ExploitParams) ToURLValues() url.Values {
	values := make(url.Values)
	values.Add("query", p.Query.String())

	if len(p.Facets) > 0 {
		values.Add("facets", strings.Join(p.Facets, ","))
	}

	if p.Page > 1 {
		values.Add("page", fmt.Sprint(p.Page))
	}

	return values
}
