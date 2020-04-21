# Shodan API for Golang

[![GoDoc](https://img.shields.io/badge/docs-pkg.go.dev-informational?logo=go&style=flat-square)](https://pkg.go.dev/github.com/shadowscatcher/shodan?tab=doc)
![Build](https://img.shields.io/github/workflow/status/shadowscatcher/shodan/Go?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/shadowscatcher/shodan?style=flat-square)](https://goreportcard.com/report/github.com/shadowscatcher/shodan)
[![MIT License](https://img.shields.io/badge/license-MIT-informational.svg?style=flat-square)](LICENSE)

Yet another one Golang implementation of Shodan REST API client. 
This library inspired by [Nikita Safonov](https://github.com/ns3777k/)'s [go-shodan library](https://github.com/ns3777k/go-shodan), but has different data models and query syntax.

## Features

- Library intended to be the most comprehensive and documented out there, letting you learn about all the API methods, search filters and gathered data types. The documentation is a work in progress.
- Search syntax allows you to change query without string formatting:

```go
package main

import (
	"context"
	"github.com/shadowscatcher/shodan"
	"github.com/shadowscatcher/shodan/search"
	"github.com/shadowscatcher/shodan/search/ssl_versions"
	"log"
	"net/http"
	"os"
)

func main() {
	nginxSearch := search.Params{
		Page:1,
		Query: search.Query{
			Product: "nginx",
			ASN:  "AS14618",
			SSLOpts: search.SSLOpts{
				Cert: search.CertOptions{
					Expired: true,
				},
				Version: ssl_versions.TLSv1_2,
			},
		},
	}

	client, _ := shodan.GetClient(os.Getenv("SHODAN_API_KEY"), http.DefaultClient, true)
	ctx := context.Background()
	result, err := client.Search(ctx, nginxSearch)
	if err != nil {
		log.Fatal(err)
	}

	for _, match := range result.Matches {
		// a lot of returned data can be used in another searches
		// it's easy because you will get response with almost all possible fields, just don't forget to check them
		if match.HTTP != nil && match.HTTP.Favicon != nil {
			//newQuery := search.Query{HTTP: search.HTTP{Favicon: search.Favicon{Hash: match.HTTP.Favicon.Hash}}}
		}
	}
	
	// later on you can change every part of search query or parameters:
	nginxSearch.Page++  // for example, increase page
	nginxSearch.Query.Port = 443 // or add new search term
	result, err = client.Search(ctx, nginxSearch)  // and reuse modified parameters object
	if err != nil {
		log.Fatal(err)
	}
}
```

- Search results contain a lot of types that are ignored by most of the existing libraries, documented where possible:

```go
for _, match := range result.Matches {
	if match.MongoDB != nil && !match.MongoDB.Authentication {
		fmt.Println("exposed mongodb:", match.IpAndPort())
		databases := match.MongoDB.ListDatabases.Databases

		fmt.Println("databases:", len(databases), "size:", match.MongoDB.ListDatabases.TotalSize)
		for _, db := range databases {
			for _, collectionName := range db.Collections {
				fmt.Println(collectionName)
			}
		}
	}
		
	if match.SSL != nil && match.SSL.Cert.Expired {
		fmt.Println("expired certificate:", match.IpAndPort())
	}
		
	if match.Elastic != nil {
		fmt.Println("exposed elastic:", match.IpAndPort())
		for indexName, index := range match.Elastic.Indices {
			fmt.Println(indexName, index.UUID)
		}
	}
}
```

 - The client can be configured to automatically make one second pause between requests (this interval required by Shodan's API terms of service).