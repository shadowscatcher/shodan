package services

type DNS struct {
	//  Whether or not the server allows recursive lookups
	Recursive bool `json:"recursive"`

	//  Name of the resolver used to process the query
	ResolverHostname *string `json:"resolver_hostname"`

	//  Unique identifier of the resolver
	ResolverID *string `json:"resolver_id"`

	//  Name of the DNS software
	Software *string `json:"software"`
}
