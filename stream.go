package shodan

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shadowscatcher/shodan/models"
	"github.com/shadowscatcher/shodan/routes"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type StreamClient struct {
	apiKey string
	mu     *sync.Mutex
	HTTP   *http.Client
}

func GetStreamClient(key string, client *http.Client) (*StreamClient, error) {
	if key == "" {
		return nil, errors.New("empty API key")
	}

	if client == nil {
		return nil, errors.New("HTTP client is nil")
	}

	return &StreamClient{
		apiKey: key,
		mu:     &sync.Mutex{},
		HTTP:   client,
	}, nil
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

	return req.WithContext(ctx), nil
}

func (s *StreamClient) makeRequest(ctx context.Context, route string) (response *http.Response, err error) {
	fullUrl := routes.ApiStream + route
	uri, err := url.Parse(fullUrl)

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
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errFromResponse(response)
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

// This stream provides ALL of the data that Shodan collects. Use this stream if you need access to everything and/or
// want to store your own Shodan database locally. If you only care about specific ports, please use the Ports stream
func (s *StreamClient) Banners(ctx context.Context) (chan models.Service, error) {
	return s.subscribe(ctx, routes.ShodanBanners)
}

// This stream provides a filtered, bandwidth-saving view of the Banners stream in case you are only interested in
// devices located in certain ASNs
func (s *StreamClient) ASN(ctx context.Context, asns []string) (chan models.Service, error) {
	route := fmt.Sprintf(routes.ShodanAsn, strings.Join(asns, ","))
	return s.subscribe(ctx, route)
}

// This stream provides a filtered, bandwidth-saving view of the Banners stream in case you are only interested in
// devices located in certain countries
func (s *StreamClient) Countries(ctx context.Context, countries []string) (chan models.Service, error) {
	route := fmt.Sprintf(routes.ShodanCountries, strings.Join(countries, ","))
	return s.subscribe(ctx, route)
}

// Only returns banner data for the list of specified ports. This stream provides a filtered, bandwidth-saving view of
// the Banners stream in case you are only interested in a specific list of ports.
func (s *StreamClient) Ports(ctx context.Context, ports []int) (chan models.Service, error) {
	portList := make([]string, len(ports), len(ports))

	for i, port := range ports {
		portList[i] = fmt.Sprint(port)
	}

	route := fmt.Sprintf(routes.ShodanPortsList, strings.Join(portList, ","))
	return s.subscribe(ctx, route)
}

// Subscribe to banners discovered on all IP ranges described in the network alert
func (s *StreamClient) Alerts(ctx context.Context) (chan models.Service, error) {
	return s.subscribe(ctx, routes.ShodanAlerts)
}

// Subscribe to banners discovered on the IP range defined in a specific network alert
func (s *StreamClient) Alert(ctx context.Context, alertId string) (chan models.Service, error) {
	route := fmt.Sprintf(routes.ShodanAlertId, alertId)
	return s.subscribe(ctx, route)
}
