package search

import (
	"fmt"
	"github.com/shadowscatcher/shodan/search/link_types"
	"github.com/shadowscatcher/shodan/search/ssl_versions"
	"reflect"
	"strings"
)

type Query struct {
	// Raw query text. You can use only this field, if you want to
	Text string

	// Any text
	All string `shodan_search:"all"`

	// IP-address
	IP string `shodan_search:"ip"`

	// Only show results that were collected after the given date (dd/mm/yyyy).
	After string `shodan_search:"after"`

	// The Autonomous System Number that identifies the network the device is on; ex: "AS15169"
	ASN string `shodan_search:"asn"`

	//O nly show results that were collected before the given date (dd/mm/yyyy).
	Before string `shodan_search:"before"`

	// Show results that are located in the given city.
	City string `shodan_search:"city"`

	// Show results that are located within the given country.
	Country string `shodan_search:"country"`

	// Common platform enumeration
	CPE string `shodan_search:"cpe"`

	// Device type; ex: "printer", "router"
	Device string `shodan_search:"device"`

	// There are 2 modes to the geo filter: radius and bounding box.
	// To limit results based on a radius around a pair of latitude/longitude, provide 3 parameters; ex: geo:50,50,100.
	// If you want to find all results within a bounding box, supply the top left and bottom right coordinates for
	// the region; ex: geo:10,10,50,50.
	Geo string `shodan_search:"geo"`

	// Search for hosts that contain the given value in their hostname
	Hostname string `shodan_search:"hostname"`

	// Find devices based on the upstream owner of the IP netblock
	ISP string `shodan_search:"isp"`

	// Find devices depending on their connection to the Internet
	Link link_types.LinkType `shodan_search:"link"`

	// Search by netblock using CIDR notation; ex: net:69.84.207.0/24
	Net string `shodan_search:"net"`

	// Find devices based on the owner of the IP netblock.
	Org string `shodan_search:"org"`

	// Filter results based on the operating system of the device.
	OS string `shodan_search:"os"`

	// Search by postal code.
	Postal string `shodan_search:"postal"`

	// Filter using the name of the software/product; ex: product:Apache
	Product string `shodan_search:"product"`

	// Search for devices based on the state/region they are located in
	State string `shodan_search:"state"`

	// Filter the results to include only products of the given version; ex: product:apache version:1.3.37
	Version string `shodan_search:"version"`

	// Search all SSL data
	SSL string `shodan_search:"ssl"`

	// Find devices based on the services/ports that are publicly exposed on the Internet
	Port int `shodan_search:"port"`

	// Hash of the "data" property
	Hash int `shodan_search:"hash"`

	// If "true" only show results that were discovered on IPv6
	HasIPV6 bool `shodan_search:"has_ipv6"`

	// If "true" only show results that have a screenshot available
	HasScreenshot bool `shodan_search:"has_screenshot"`

	// If "true" only show results that have SSL
	HasSSL bool `shodan_search:"has_ssl"`

	// If "true" only show results that have vulnerabilities. Enterpise only.
	HasVuln bool `shodan_search:"has_vuln"`

	// Region code
	Region int `shodan_search:"region"`

	// Host tag. Enterprise only.
	Tag string `shodan_search:"tag"`

	// signature unknown
	Scan string `shodan_search:"scan"`

	// filter by vulnerability. Only available to academic users or Small Business API subscription and higher.
	Vuln string `shodan_search:"vuln"`

	SSLOpts    `shodan_search:"ssl"`
	Bitcoin    `shodan_search:"bitcoin"`
	Telnet     `shodan_search:"telnet"`
	HTTP       `shodan_search:"http"`
	NTP        `shodan_search:"ntp"`
	Screenshot `shodan_search:"screenshot"`
	Shodan     `shodan_search:"shodan"`
	SNMP       `shodan_search:"snmp"`
	SSH        `shodan_search:"ssh"`
}

type SSLOpts struct {
	// Number of certificates in the chain
	ChainCount int `shodan_search:"chain_count"`

	// Application layer protocols such as HTTP/2 ("h2")
	ALPN string `shodan_search:"alpn"`

	// Possible values: SSLv2, SSLv3, TLSv1, TLSv1.1, TLSv1.2
	Version ssl_versions.SSLVersion `shodan_search:"version"`

	// Various certificate options
	Cert   CertOptions `shodan_search:"cert"`
	Cipher Cipher      `shodan_search:"cipher"`
}

type Pubkey struct {
	// Number of bits in the public key
	Bits int `shodan_search:"bits"`

	// Public key type
	Type string `shodan_search:"type"`
}

type Cipher struct {
	// SSL version of the preferred cipher
	Version string `shodan_search:"version"`

	// Number of bits in the preferred cipher
	Bits int `shodan_search:"bits"`

	// Name of the preferred cipher
	Name string `shodan_search:"name"`
}

type CertOptions struct {
	// Certificate algorithm
	Alg string `shodan_search:"alg"`

	// Whether the SSL certificate is expired or not
	Expired bool `shodan_search:"expired"`

	// Names of extensions in the certificate
	Extension string `shodan_search:"extension"`

	// Serial number as string
	Serial string `shodan_search:"serial"`

	// SHA-1 fingerprint
	Fingerprint string `shodan_search:"fingerprint"`

	// Cert issuer options
	Issuer CertEntity `shodan_search:"issuer"`

	// Cert subject options
	Subject CertEntity `shodan_search:"subject"`

	Pubkey Pubkey `shodan_search:"pubkey"`
}

type CertEntity struct {
	// Common name
	CN string `shodan_search:"cn"`
}

type Bitcoin struct {
	//Find Bitcoin servers that had the given IP in their list of peers
	IP string `shodan_search:"ip"`

	//Find Bitcoin servers that return the given number of IPs in the list of peers
	IPCount int `shodan_search:"ip_count"`

	//Find Bitcoin servers that had IPs with the given port in their list of peers
	Port int `shodan_search:"port"`

	//Filter results based on the Bitcoin protocol version
	Version string `shodan_search:"version"`
}

type Telnet struct {
	// Search all the options
	Option string `shodan_search:"option"`

	// The server requests the client to support these options
	Do string `shodan_search:"do"`

	// The server requests the client to not support these options
	Dont string `shodan_search:"dont"`

	// The server supports these options
	Will string `shodan_search:"will"`

	// The server doesnt support these options
	Wont string `shodan_search:"wont"`
}

type HTTP struct {
	// Name of web technology used on the website
	Component string `shodan_search:"component"`

	// Category of web components used on the website
	ComponentCategory string `shodan_search:"component_category"`

	// Search the HTML of the website for the given value.
	HTML string `shodan_search:"html"`

	// Search the title of the website
	Title string `shodan_search:"title"`

	// Response status code
	Status int `shodan_search:"status"`

	// Hash of the website HTML
	HTMLHash int `shodan_search:"html_hash"`

	// Hash of website favicon.ico file
	Favicon Favicon `shodan_search:"favicon"`

	// Hash of website robots.txt file
	RobotsHash int `shodan_search:"robots_hash"`

	// Search in contents of website's security.txt
	SecurityTxt string `shodan_search:"securitytxt"`

	// Search by Web Application Firewall vendor/name
	WAF string `shodan_search:"waf"`
}

type NTP struct {
	// Find NTP servers that had the given IP in their monlist.
	IP string `shodan_search:"ip"`

	// Find NTP servers that return the given number of IPs in the initial monlist response.
	IPCount int `shodan_search:"ip_count"`

	// Find NTP servers that had IPs with the given port in their monlist.
	Port int `shodan_search:"port"`

	// Whether or not more IPs were available for the given NTP server.
	More bool `shodan_search:"more"`
}

type Screenshot struct {
	// Label of screenshot (kind of tag, like "login", "windows")
	Label string `shodan_search:"label"`
}

type Favicon struct {
	// Hash of website favicon.ico file
	Hash int `shodan_search:"hash"`
}

type Shodan struct {
	// Filter by shodan crawler module
	Module string `shodan_search:"module"`
}

type SNMP struct {
	// SNMP contact address
	Contact string `shodan_search:"contact"`

	// SNMP server name
	Name string `shodan_search:"name"`

	Location string `shodan_search:"location"`
}

type SSH struct {
	// HASSH Md5 fingerprint hash
	HASSH string `shodan_search:"hassh"`

	Type string `shodan_search:"type"`
}

func (s *Query) String() string {
	marshaled := marshalQueryParam(s, true, "")
	switch {
	case marshaled == "":
		return s.Text
	case s.Text == "":
		return marshaled
	}

	return fmt.Sprintf("%s %s", s.Text, marshaled)
}

func marshalQueryParamField(typeField reflect.StructField, valueField reflect.Value, tag string) string {
	switch typeField.Type {
	case stringType, sslVersionType, linkTypeType, explotTypeType:
		value := valueField.String()
		if value != "" {
			if strings.Contains(value, " ") {
				value = fmt.Sprintf("\"%s\"", value)
			}
			return fmt.Sprintf("%s:%s ", tag, value)
		}
	case intType:
		value := valueField.Int()
		if value != 0 {
			return fmt.Sprintf("%s:%d ", tag, value)
		}
	case boolType:
		value := valueField.Bool()
		if value {
			return fmt.Sprintf("%s:true ", tag)
		}
	default:
		value := valueField.Interface()
		return marshalQueryParam(value, false, tag)
	}

	return ""
}

func marshalQueryParam(v interface{}, concreteType bool, topLevelTag string) string {
	var object reflect.Value
	if concreteType {
		object = reflect.ValueOf(v).Elem()
	} else {
		object = reflect.ValueOf(v)
	}
	var result string

	for i := 0; i < object.NumField(); i++ {
		valueField := object.Field(i)
		typeField := object.Type().Field(i)
		tag := typeField.Tag
		searchTag := tag.Get("shodan_search")
		if searchTag == "" {
			continue
		}

		if topLevelTag != "" {
			searchTag = fmt.Sprintf("%s.%s", topLevelTag, searchTag)
		}

		result += marshalQueryParamField(typeField, valueField, searchTag)
	}

	return strings.Trim(result, " ")
}
