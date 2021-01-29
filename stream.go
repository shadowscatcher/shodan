package shodan

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/shadowscatcher/shodan/models"
	"github.com/shadowscatcher/shodan/routes"
)

// StreamClient is a client with all stream-related methods. Use GetStreamClient to create instance
type StreamClient struct {
	apiKey string
	HTTP   *http.Client
	// StreamResponseHook allows you to intercept response before stream reading. If it returns an error, method exits
	ResponseHook func(response *http.Response, err error) error
}

// GetStreamClient creates StreamClient instance. If you want to use a proxy, configure http.Client.
func GetStreamClient(key string, client *http.Client) (*StreamClient, error) {
	if key == "" {
		return nil, errors.New("empty API key")
	}

	if client == nil {
		return nil, errors.New("HTTP client is nil")
	}

	return &StreamClient{
		apiKey:       key,
		HTTP:         client,
		ResponseHook: defaultResponseHook,
	}, nil
}

// Banners stream provides ALL of the data that Shodan collects. Use this stream if you need access to everything and/or
// want to store your own Shodan database locally. If you only care about specific ports, please use the Ports stream
func (s *StreamClient) Banners(ctx context.Context) (chan models.Service, error) {
	return s.subscribe(ctx, routes.ShodanBanners)
}

// ASN stream provides a filtered, bandwidth-saving view of the Banners stream in case you are only interested in
// devices located in certain ASNs
func (s *StreamClient) ASN(ctx context.Context, asns []string) (chan models.Service, error) {
	if asns == nil || len(asns) == 0 {
		return nil, errors.New("asns are required")
	}

	route := fmt.Sprintf(routes.ShodanAsn, strings.Join(asns, ","))
	return s.subscribe(ctx, route)
}

// Countries stream provides a filtered, bandwidth-saving view of the Banners stream in case you are only interested in
// devices located in certain countries
func (s *StreamClient) Countries(ctx context.Context, countries []string) (chan models.Service, error) {
	if countries == nil || len(countries) == 0 {
		return nil, errors.New("countries are required")
	}

	route := fmt.Sprintf(routes.ShodanCountries, strings.Join(countries, ","))
	return s.subscribe(ctx, route)
}

// Ports stream only returns banner data for the list of specified ports. This stream provides a filtered, bandwidth-saving view of
// the Banners stream in case you are only interested in a specific list of ports.
func (s *StreamClient) Ports(ctx context.Context, ports []int) (chan models.Service, error) {
	if ports == nil || len(ports) == 0 {
		return nil, errors.New("ports are required")
	}

	portList := make([]string, len(ports), len(ports))

	for i, port := range ports {
		portList[i] = fmt.Sprint(port)
	}

	route := fmt.Sprintf(routes.ShodanPortsList, strings.Join(portList, ","))
	return s.subscribe(ctx, route)
}

// Alerts stream allows to subscribe to banners discovered on all IP ranges described in the network alert
func (s *StreamClient) Alerts(ctx context.Context) (chan models.Service, error) {
	return s.subscribe(ctx, routes.ShodanAlert)
}

// Alert stream allows to subscribe to banners discovered on the IP range defined in a specific network alert
func (s *StreamClient) Alert(ctx context.Context, alertID string) (chan models.Service, error) {
	if alertID == "" {
		return nil, errors.New("alertID is required")
	}

	route := fmt.Sprintf(routes.ShodanAlertId, alertID)
	return s.subscribe(ctx, route)
}

// Tags is a filtered version of the "banners" stream to only return banners that match the tags of interest
func (s *StreamClient) Tags(ctx context.Context, tags []string) (chan models.Service, error) {
	if tags == nil || len(tags) == 0 {
		return nil, errors.New("tags are required")
	}

	route := fmt.Sprintf(routes.ShodanTags, strings.Join(tags, ","))
	return s.subscribe(ctx, route)
}

// Vulns is a filtered version of the "banners" stream to only return banners that match the vulnerabilities of interest
func (s *StreamClient) Vulns(ctx context.Context, vulns []string) (chan models.Service, error) {
	if vulns == nil || len(vulns) == 0 {
		return nil, errors.New("vulns are required")
	}

	route := fmt.Sprintf(routes.ShodanVulns, strings.Join(vulns, ","))
	return s.subscribe(ctx, route)
}

func (s *StreamClient) defaultParams() url.Values {
	params := make(url.Values)
	params.Set("key", s.apiKey)
	return params
}

func (s *StreamClient) createRequest(
	ctx context.Context,
	uri *url.URL) (*http.Request, error) {

	params := s.defaultParams()
	uri.RawQuery = params.Encode()

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

func (s *StreamClient) makeRequest(ctx context.Context, route string) (response *http.Response, err error) {
	fullURL := routes.ApiStream + route
	uri, err := url.Parse(fullURL)

	if err != nil {
		return
	}

	request, err := s.createRequest(ctx, uri)
	if err != nil {
		return
	}

	response, err = s.HTTP.Do(request)
	return
}

func (s *StreamClient) subscribe(ctx context.Context, route string) (chan models.Service, error) {
	resultChan := make(chan models.Service)
	response, err := s.makeRequest(ctx, route)

	if err = s.ResponseHook(response, err); err != nil {
		return nil, err
	}

	go func(c chan models.Service, body io.ReadCloser) {
		defer body.Close()
		reader := bufio.NewReader(body)
		for {
			var service models.Service
			chunk, err := reader.ReadBytes('\n')

			if err != nil {
				close(c)
				break
			}

			trimmed := bytes.TrimRight(chunk, "\r\n")
			if len(trimmed) == 0 {
				continue
			}

			err = json.Unmarshal(trimmed, &service)

			if err != nil {
				close(c)
				break
			}

			c <- service
		}
	}(resultChan, response.Body)

	return resultChan, nil
}

// Default response hook only checks for status != 200
func defaultResponseHook(response *http.Response, err error) error {
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errFromResponse(response)
	}

	return nil
}
