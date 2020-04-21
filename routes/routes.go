package routes

const (
	ApiRoot     = "https://api.shodan.io"
	ApiExploits = "https://exploits.shodan.io/api"
	ApiStream   = "https://stream.shodan.io"

	Search = "/search"
	Count  = "/count"

	ShodanHostView                       = "/shodan/host/%s"
	ShodanHostCount                      = "/shodan/host/count"
	ShodanHostSearch                     = "/shodan/host/search"
	ShodanHostSearchTokens               = "/shodan/host/search/tokens"
	ShodanHostSearchFilters              = "/shodan/host/search/filters"
	ShodanHostSearchFacets               = "/shodan/host/search/facets"
	ShodanPorts                          = "/shodan/ports"
	ShodanServices                       = "/shodan/services"
	ShodanProtocols                      = "/shodan/protocols"
	ShodanScan                           = "/shodan/scan"
	ShodanScans                          = "/shodan/scans"
	ShodanScanView                       = "/shodan/scan/%s"
	ShodanScanInternet                   = "/shodan/scan/internet"
	ShodanQuery                          = "/shodan/query"
	ShodanQuerySearch                    = "/shodan/query/search"
	ShodanQueryTags                      = "/shodan/query/tags"
	ShodanAlert                          = "/shodan/alert"
	ShodanAlertInfo                      = "/shodan/alert/info"
	ShodanAlertId                        = "/shodan/alert/%s"
	ShodanAlertNotifierId                = "/shodan/alert/%s/notifier/%s"
	ShodanAlertIdInfo                    = "/shodan/alert/%s/info"
	ShodanAlertTriggers                  = "/shodan/alert/triggers"                // Return a list of available triggers that can be enabled for alerts
	ShodanAlertTriggerAction             = "/shodan/alert/%s/trigger/%s"           // Enable/disable the given trigger on the alert (HTTP put/delete) [alertId, trigger]
	ShodanAlertTriggerNotificationAction = "/shodan/alert/%s/trigger/%s/ignore/%s" // Ignore/enable trigger notifications for the provided IP and port [alertId, trigger, ip:port]
	ShodanData                           = "/shodan/data"
	ShodanDataset                        = "/shodan/data/%s"
	Org                                  = "/org"
	OrgMember                            = "/org/member/%s"
	AccountProfile                       = "/account/profile"
	DnsResolve                           = "/dns/resolve"
	DnsReverse                           = "/dns/reverse"
	DnsDomain                            = "/dns/domain/%s"
	ToolsHTTPHeaders                     = "/tools/httpheaders"
	ToolsMyIP                            = "/tools/myip"
	ApiInfo                              = "/api-info"
	LabsHoneyscore                       = "/labs/honeyscore/%s"

	// stream API routes
	ShodanBanners   = "/shodan/banners"
	ShodanAsn       = "/shodan/asn/%s"
	ShodanCountries = "/shodan/countries/%s"
	ShodanPortsList = "/shodan/ports/%s"
	ShodanAlerts    = "/shodan/alerts"
	ShodanTags      = "/shodan/tags/%s"
	ShodanVulns     = "/shodan/vulns/%s"

	// notifier API routes
	Notifier         = "/notifier"
	NotifierId       = "/notifier/%s"
	NotifierProvider = "/notifier/provider"
)
