# Shodan API for Golang

Yet another one Golang implementation of Shodan REST API client. 
This library was inspired by [Nikita Safonov](https://github.com/ns3777k/)'s [go-shodan library](https://github.com/ns3777k/go-shodan), but has different data models and query syntax to suit my own needs.

## Features

- Library is intended to be the most comprehensive and documented out there, letting you learn about all of the API features and gathered data types. The documentation is a work in progress.
- Search syntax allows you to change query without without string formatting:

```go
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
// later on you can change every part of search query or parameters:
nginxSearch.Page++  // for example, increase page
nginxSearch.Query.Port = 443 // or add new search term 
result2, err := client.Search(ctx, nginxSearch)
```

- Search results contains a lot of types that are ignored by most of the existing libraries, documented where possible:

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

 - The client can be configured to automatically make one second pause between requests (this interval is required by Shodan's API terms of service).