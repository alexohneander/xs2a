package xs2a

import (
	"fmt"
	"net/http"
	"time"
)

type (
	Client struct {
		host          string
		httpClient    *http.Client
		clientId      string
		codeChallenge string
	}
)

func NewClient(host string, clientId, codeChallenge string, timeout time.Duration) *Client {
	client := &http.Client{
		Timeout: timeout,
	}
	return &Client{
		host:          host,
		httpClient:    client,
		clientId:      clientId,
		codeChallenge: codeChallenge,
	}
}

func (c *Client) do(method, endpoint string, params map[string]string) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s/%s", c.host, endpoint)
	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {
		return nil, err
	}
	// req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	for key, val := range params {
		q.Set(key, val)
	}
	req.URL.RawQuery = q.Encode()
	return c.httpClient.Do(req)
}
