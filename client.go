package shodan

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/shadowscatcher/shodan/models"
	"github.com/shadowscatcher/shodan/routes"

	"sync"
	"time"
)

const intSecond = int64(time.Second)

// Client is a type with all non-stream methods. Use GetClient to create instance
type Client struct {
	apiKey          string
	waitFunc        func(*Client)
	lastRequestTime int64
	mu              *sync.Mutex

	HTTP *http.Client
}

// All API plans are subject to a rate limit of 1 request per second. See https://developer.shodan.io/billing/signup
func waitASecond(client *Client) {
	delta := time.Now().UnixNano() - client.lastRequestTime
	if delta < intSecond {
		time.Sleep(time.Duration(intSecond - delta))
	}
}

// GetClient creates Client instance. apiKey is required to work with API. If you want to use a proxy, configure http.Client.
// If you need to disable throttling, set wait to false
func GetClient(apiKey string, client *http.Client, wait bool) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("empty API key")
	}

	if client == nil {
		return nil, errors.New("HTTP client is nil")
	}

	shodanClient := &Client{
		apiKey:          apiKey,
		lastRequestTime: 0,
		mu:              &sync.Mutex{},

		HTTP: client,
	}

	if wait {
		shodanClient.waitFunc = waitASecond
	} else {
		shodanClient.waitFunc = func(*Client) {}
	}

	return shodanClient, nil
}

// parameters applicable to all requests
func (c *Client) defaultParams() url.Values {
	vals := make(url.Values)
	vals.Set("key", c.apiKey)
	return vals
}

// performs actual request to Shodan; waits request timeout, if client is instructed to do so
func (c *Client) do(r *http.Request) (response *http.Response, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.waitFunc(c)
	c.lastRequestTime = time.Now().UnixNano()

	return c.HTTP.Do(r)
}

// ensures that every request have api key in parameters
func (c *Client) ensureRequestParams(params url.Values) url.Values {
	if params == nil || len(params) == 0 {
		return c.defaultParams()
	} else if params.Get("key") == "" {
		params.Set("key", c.apiKey)
	}
	return params
}

// extract Shodan error from response body
func errFromResponse(response *http.Response) error {
	var errResp models.Error
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("can't read error response body: %s", err)
	}

	err = json.Unmarshal(content, &errResp)

	if err != nil {
		return errors.New(string(content))
	}

	return errors.New(errResp.Error)
}

// creates HTTP request with all required parameters
func (c *Client) createRequest(
	ctx context.Context,
	method string,
	uri *url.URL,
	params url.Values,
	body io.Reader,
	header http.Header) (*http.Request, error) {

	params = c.ensureRequestParams(params)
	uri.RawQuery = params.Encode()

	req, err := http.NewRequest(method, uri.String(), body)
	if err != nil {
		return nil, err
	}

	if header != nil {
		req.Header = header
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// creates HTTP request for root api (https://api.shodan.io) method
func (c *Client) createRootRequest(
	ctx context.Context,
	method, endpoint string,
	params url.Values,
	body io.Reader,
	header http.Header) (*http.Request, error) {

	uri, err := url.Parse(routes.ApiRoot + endpoint)
	if err != nil {
		return nil, err
	}

	return c.createRequest(ctx, method, uri, params, body, header)
}

// creates HTTP request for exploits api (https://exploits.shodan.io/api) method
func (c *Client) createExploitRequest(
	ctx context.Context,
	method, endpoint string,
	params url.Values,
	body io.Reader,
	header http.Header) (*http.Request, error) {

	uri, err := url.Parse(routes.ApiExploits + endpoint)
	if err != nil {
		return nil, err
	}

	return c.createRequest(ctx, method, uri, params, body, header)
}

// reads or unmarshals HTTP response
func (c *Client) readResponse(to interface{}, body io.Reader) error {
	var err error

	if w, ok := to.(io.Writer); ok {
		_, err = io.Copy(w, body)
	} else {
		decoder := json.NewDecoder(body)
		err = decoder.Decode(to)
	}

	return err
}

func (c *Client) requestAndRead(req *http.Request, result interface{}) (err error) {
	response, err := c.do(req)
	if err != nil {
		return
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		err = errFromResponse(response)
		return
	}

	err = c.readResponse(result, response.Body)
	return
}

// composes request and proceeds with it; Unmarshals results
func (c *Client) request(
	ctx context.Context,
	method, route string,
	params url.Values,
	body io.Reader,
	header http.Header,
	result interface{}) (err error) {

	req, err := c.createRootRequest(ctx, method, route, params, body, header)
	if err != nil {
		return
	}

	err = c.requestAndRead(req, result)
	return
}

func (c *Client) requestExploits(
	ctx context.Context,
	method, route string,
	params url.Values,
	body io.Reader,
	header http.Header,
	result interface{}) (err error) {

	req, err := c.createExploitRequest(ctx, method, route, params, body, header)
	if err != nil {
		return
	}

	err = c.requestAndRead(req, result)
	return
}

// most endpoints are GET endpoints of API root
func (c *Client) get(
	ctx context.Context,
	route string,
	params url.Values,
	result interface{}) (err error) {

	err = c.request(ctx, http.MethodGet, route, params, nil, nil, result)
	return
}
