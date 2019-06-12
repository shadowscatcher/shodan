package search

import (
	"github.com/shadowscatcher/shodan/search/exploit_types"
	"github.com/shadowscatcher/shodan/search/link_types"
	"github.com/shadowscatcher/shodan/search/ssl_versions"
	"reflect"
)

// types for marshaling
var stringType = reflect.TypeOf("")
var intType = reflect.TypeOf(0)
var boolType = reflect.TypeOf(true)
var sslVersionType = reflect.TypeOf((*ssl_versions.SSLVersion)(nil)).Elem()
var linkTypeType = reflect.TypeOf((*link_types.LinkType)(nil)).Elem()
var explotTypeType = reflect.TypeOf((*exploit_types.ExploitType)(nil)).Elem()
