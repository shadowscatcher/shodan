package shodan

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shadowscatcher/shodan/models"
	"github.com/shadowscatcher/shodan/routes"
	"github.com/shadowscatcher/shodan/search"
	"net/http"
	"net/url"
	"strings"
)

var emptyParams = errors.New("empty parameters")
var emptyAlertId = errors.New("empty alert id")
var emptyTriggerName = errors.New("empty trigger name")
var emptyService = errors.New("empty service")
var emptyUsername = errors.New("empty username")
var bigRequest = errors.New("request is too big")

const (
	hostnamesLenLimit = 3575
	ipsLenLimit       = 3369
)

// Returns all services that have been found on the given host IP
func (c *Client) Host(ctx context.Context, params search.HostParams) (result models.Host, err error) {
	route := fmt.Sprintf(routes.ShodanHostView, params.IP)
	err = c.get(ctx, route, params.ToURLValues(), &result)
	return
}

// Search Shodan without Results
// This method behaves identical to Search() with the only difference that this method does not return any host results,
// it only returns the total number of results that matched the query and any facet information that was requested.
// As a result this method does not consume query credits.
func (c *Client) Count(ctx context.Context, params search.Params) (result models.SearchResult, err error) {
	err = c.get(ctx, routes.ShodanHostCount, params.ToURLValues(), &result)
	return
}

// Search using the same query syntax as the website and use facets to get summary information for different properties
// This method may use API query credits depending on usage.
// If any of the following criteria are met, your account will be deducated 1 query credit:
// * The search query contains a filter.
// * Accessing results past the 1st page using the "page". For every 100 results past the 1st page 1 query credit
// is deducted.
func (c *Client) Search(ctx context.Context, params search.Params) (result models.SearchResult, err error) {
	// todo: check: minify seems ignored
	values := params.ToURLValues()
	if len(values) == 0 {
		err = emptyParams
		return
	}

	err = c.get(ctx, routes.ShodanHostSearch, values, &result)
	return
}

// Break the search query into tokens
// This method lets you determine which filters are being used by the query string and what parameters were provided
// to the filters.
func (c *Client) SearchTokens(ctx context.Context, params search.Params) (result models.Tokens, err error) {
	values := params.ToURLValues()
	if len(values) == 0 {
		err = emptyParams
		return
	}

	err = c.get(ctx, routes.ShodanHostSearchTokens, values, &result)
	return
}

// This method returns a list of port numbers that the crawlers are looking for
func (c *Client) Ports(ctx context.Context) (result []int, err error) {
	err = c.get(ctx, routes.ShodanPorts, nil, &result)
	return
}

// This method returns a map containing all the protocols that can be used when launching an Internet scan
func (c *Client) Protocols(ctx context.Context) (result map[string]string, err error) {
	err = c.get(ctx, routes.ShodanProtocols, nil, &result)
	return
}

// This method returns a map containing all the services Shodan can detect
func (c *Client) Services(ctx context.Context) (result map[string]string, err error) {
	err = c.get(ctx, routes.ShodanServices, nil, &result)
	return
}

// Request Shodan to crawl an IP/netblock
// This method uses API scan credits: 1 IP consumes 1 scan credit. You must have a paid API plan
// (either one-time payment or subscription) in order to use this method
func (c *Client) SubmitScan(ctx context.Context, ips []string, force bool) (result models.Scan, err error) {
	if ips == nil || len(ips) == 0 {
		err = emptyParams
		return
	}

	params := make(url.Values)
	params.Set("ips", strings.Join(ips, ","))
	if force {
		params.Set("force", "true")
	}

	body := strings.NewReader(params.Encode())
	header := make(http.Header)
	header.Set("Content-Type", "application/x-www-form-urlencoded")

	err = c.request(ctx, http.MethodPost, routes.ShodanScan, nil, body, header, &result)
	return
}

// Returns a list of all your scans
func (c *Client) ListScans(ctx context.Context, page uint) (result models.ScanList, err error) {
	params := make(url.Values)
	if page < 1 {
		params.Set("page", "1")
	} else {
		params.Set("page", fmt.Sprint(page))
	}
	err = c.get(ctx, routes.ShodanScans, params, &result)
	return
}

// Check the progress of a previously submitted scan request
func (c *Client) GetScan(ctx context.Context, scanId string) (result models.Scan, err error) {
	if scanId == "" {
		err = errors.New("empty scanId")
		return
	}

	route := fmt.Sprintf(routes.ShodanScanView, scanId)
	err = c.request(ctx, http.MethodGet, route, nil, nil, nil, &result)
	return
}

// Use this method to request Shodan to crawl the Internet for a specific port.
// This method is restricted to security researchers and companies with a Shodan Enterprise Data license. To apply
// for access to this method as a researcher, please email jmath@shodan.io with information about your project.
// Access is restricted to prevent abuse.
func (c *Client) ScanInternet(ctx context.Context, port uint16, protocol string) (result models.Scan, err error) {
	if protocol == "" {
		err = errors.New("empty protocol")
		return
	}

	params := make(url.Values)
	params.Set("port", fmt.Sprint(port))
	params.Set("protocol", protocol)

	body := strings.NewReader(params.Encode())
	header := make(http.Header)
	header.Set("Content-Type", "application/x-www-form-urlencoded")

	err = c.request(ctx, http.MethodPost, routes.ShodanScanInternet, nil, body, header, result)
	return
}

// Use this method to obtain a list of search queries that users have saved in Shodan
func (c *Client) QueryList(ctx context.Context, page uint, sort, order string) (result models.SearchQueries, err error) {
	params := make(url.Values)

	if page > 0 {
		params.Set("page", fmt.Sprint(page))
	}

	if sort != "" {
		params.Set("sort", sort)
	}

	err = c.get(ctx, routes.ShodanQuery, params, &result)
	return
}

// Use this method to search the directory of search queries that users have saved in Shodan
func (c *Client) QuerySearch(ctx context.Context, query string, page uint) (result models.SearchQueries, err error) {
	if query == "" {
		err = errors.New("empty search query")
		return
	}

	params := make(url.Values)
	params.Set("query", query)
	if page > 0 {
		params.Set("page", fmt.Sprint(page))
	}
	err = c.get(ctx, routes.ShodanQuerySearch, params, &result)
	return
}

// Use this method to obtain a list of popular tags for the saved search queries in Shodan
func (c *Client) QueryTags(ctx context.Context, size uint) (result models.QueryTags, err error) {
	params := make(url.Values)
	if size > 0 {
		params.Set("size", fmt.Sprint(size))
	}
	err = c.get(ctx, routes.ShodanQueryTags, params, &result)
	return
}

// Use this method to see a list of the datasets that are available for download
func (c *Client) Datasets(ctx context.Context) (result []models.Dataset, err error) {
	err = c.get(ctx, routes.ShodanData, nil, &result)
	return
}

// Get a list of files that are available for download from the provided dataset
func (c *Client) DatasetFiles(ctx context.Context, dataset string) (result []models.DatasetFile, err error) {
	if dataset == "" {
		err = errors.New("empty dataset id")
		return
	}

	route := fmt.Sprintf(routes.ShodanDataset, dataset)
	err = c.get(ctx, route, nil, &result)
	return
}

// Get information about your organization such as the list of its members, upgrades, authorized domains and more
func (c *Client) Org(ctx context.Context) (result models.Org, err error) {
	err = c.get(ctx, routes.Org, nil, &result)
	return
}

// Add a Shodan user to the organization and upgrade them
func (c *Client) AddOrgMember(ctx context.Context, username string, notify bool) (result models.SimpleResponse, err error) {
	if username == "" {
		err = emptyUsername
	}

	route := fmt.Sprintf(routes.OrgMember, username)
	params := make(url.Values)
	if notify {
		params.Set("notify", "true")
	}
	err = c.request(ctx, http.MethodPut, route, params, nil, nil, &result)
	return
}

// Remove and downgrade the provided member from the organization
func (c *Client) DeleteOrgMember(ctx context.Context, username string) (result models.SimpleResponse, err error) {
	if username == "" {
		err = emptyUsername
	}

	route := fmt.Sprintf(routes.OrgMember, username)
	err = c.request(ctx, http.MethodDelete, route, nil, nil, nil, &result)
	return
}

// Returns information about the Shodan account linked to this API key
func (c *Client) AccountProfile(ctx context.Context) (result models.Profile, err error) {
	err = c.get(ctx, routes.AccountProfile, nil, &result)
	return
}

// Look up the IP address for the provided list of hostnames
func (c *Client) DnsResolve(ctx context.Context, hostnames []string) (result map[string]string, err error) {
	if hostnames == nil || len(hostnames) == 0 {
		err = emptyParams
		return
	}

	joined := strings.Join(hostnames, ",")
	if len(joined) > hostnamesLenLimit {
		err = bigRequest
		return
	}

	params := make(url.Values)
	params.Set("hostnames", joined)
	err = c.get(ctx, routes.DnsResolve, params, &result)
	return
}

// Look up the hostnames that have been defined for the given list of IP addresses
func (c *Client) DnsReverse(ctx context.Context, ips []string) (result map[string][]string, err error) {
	if ips == nil || len(ips) == 0 {
		err = emptyParams
		return
	}

	joined := strings.Join(ips, ",")

	if len(joined) > ipsLenLimit {
		err = bigRequest
		return
	}
	params := make(url.Values)
	params.Set("ips", joined)
	err = c.get(ctx, routes.DnsReverse, params, &result)
	return
}

// Shows the HTTP headers that your client sends when connecting to a webserver
func (c *Client) HttpHeaders(ctx context.Context) (result map[string]string, err error) {
	err = c.get(ctx, routes.ToolsHTTPHeaders, nil, &result)
	return
}

// Get your current IP address as seen from the Internet
func (c *Client) MyIP(ctx context.Context) (result string, err error) {
	err = c.get(ctx, routes.ToolsMyIP, nil, &result)
	return
}

// Returns information about the API plan belonging to the given API key
func (c *Client) ApiInfo(ctx context.Context) (result models.ApiInfo, err error) {
	err = c.get(ctx, routes.ApiInfo, nil, &result)
	return
}

// Calculates a honeypot probability score ranging from 0 (not a honeypot) to 1.0 (is a honeypot)
func (c *Client) Honeyscore(ctx context.Context, ip string) (result float32, err error) {
	route := fmt.Sprintf(routes.LabsHoneyscore, ip)
	err = c.get(ctx, route, nil, &result)
	return
}

// Use this method to create a network alert for a defined IP/ netblock which can be used to subscribe
// to changes/events that are discovered within that range
func (c *Client) CreateAlert(ctx context.Context, alert models.Alert) (result models.AlertDetails, err error) {
	body, err := json.Marshal(alert)
	fmt.Println(string(body))
	if err != nil {
		return
	}

	header := make(http.Header)
	header.Set("Content-Type", "application/json")

	err = c.request(ctx, http.MethodPost, routes.ShodanAlert, nil, bytes.NewReader(body), header, &result)
	return
}

// Returns the information about a specific network alert
func (c *Client) AlertInfo(ctx context.Context, alertId string) (result models.AlertDetails, err error) {
	if alertId == "" {
		err = emptyAlertId
		return
	}
	route := fmt.Sprintf(routes.ShodanAlertIdInfo, alertId)

	err = c.get(ctx, route, nil, &result)
	return
}

// Remove the specified network alert
func (c *Client) DeleteAlert(ctx context.Context, alertId string) (result interface{}, err error) {
	if alertId == "" {
		err = emptyAlertId
		return
	}

	route := fmt.Sprintf(routes.ShodanAlertId, alertId)
	err = c.request(ctx, http.MethodDelete, route, nil, nil, nil, &result)
	return
}

// Returns a listing of all the network alerts that are currently active on the account
func (c *Client) ListAlerts(ctx context.Context) (result []models.AlertDetails, err error) {
	err = c.get(ctx, routes.ShodanAlertInfo, nil, &result)
	return
}

// Returns a list of all the triggers that can be enabled on network alerts
func (c *Client) ListTriggers(ctx context.Context) (result []models.Trigger, err error) {
	err = c.get(ctx, routes.ShodanAlertTriggers, nil, &result)
	return
}

// Get notifications when the specified trigger is met
func (c *Client) CreateAlertTrigger(ctx context.Context, alertId, triggerName string) (result models.SimpleResponse, err error) {
	if alertId == "" {
		err = emptyAlertId
		return
	}

	if triggerName == "" {
		err = emptyTriggerName
		return
	}

	route := fmt.Sprintf(routes.ShodanAlertTriggerAction, alertId, triggerName)
	err = c.request(ctx, http.MethodPut, route, nil, nil, nil, &result)
	return
}

// Stop getting notifications for the specified trigger
func (c *Client) DeleteAlertTrigger(ctx context.Context, alertId, triggerName string) (result models.SimpleResponse, err error) {
	if alertId == "" {
		err = emptyAlertId
		return
	}

	if triggerName == "" {
		err = emptyTriggerName
		return
	}

	route := fmt.Sprintf(routes.ShodanAlertTriggerAction, alertId, triggerName)
	err = c.request(ctx, http.MethodDelete, route, nil, nil, nil, &result)
	return
}

// Ignore the specified service when it is matched for the trigger
func (c *Client) CreateTriggerIgnore(ctx context.Context, alertId, triggerName, service string) (result models.SimpleResponse, err error) {
	if alertId == "" {
		err = emptyAlertId
		return
	}

	if triggerName == "" {
		err = emptyTriggerName
		return
	}

	if service == "" {
		err = emptyService
		return
	}

	route := fmt.Sprintf(routes.ShodanAlertTriggerNotificationAction, alertId, triggerName, service)
	err = c.request(ctx, http.MethodPut, route, nil, nil, nil, &result)
	return
}

// Start getting notifications again for the specified trigger
func (c *Client) DeleteTriggerIgnore(ctx context.Context, alertId, triggerName, service string) (result models.SimpleResponse, err error) {
	if alertId == "" {
		err = emptyAlertId
		return
	}

	if triggerName == "" {
		err = emptyTriggerName
		return
	}

	if service == "" {
		err = emptyService
		return
	}

	route := fmt.Sprintf(routes.ShodanAlertTriggerNotificationAction, alertId, triggerName, service)
	err = c.request(ctx, http.MethodDelete, route, nil, nil, nil, &result)
	return
}

// Search across a variety of data sources for exploits and use facets to get summary information
func (c *Client) ExploitSearch(ctx context.Context, params search.ExploitParams) (result models.ExploitResult, err error) {
	values := params.ToURLValues()
	if len(values) == 0 {
		err = emptyParams
		return
	}

	err = c.requestExploits(ctx, http.MethodGet, routes.Search, values, nil, nil, &result)
	return
}

// This method behaves identical to the "/search" method with the difference that it doesn't return any results
func (c *Client) ExploitCount(ctx context.Context, params search.ExploitParams) (result models.ExploitResult, err error) {
	values := params.ToURLValues()
	if len(values) == 0 {
		err = emptyParams
		return
	}

	err = c.requestExploits(ctx, http.MethodGet, routes.Count, values, nil, nil, &result)
	return
}
