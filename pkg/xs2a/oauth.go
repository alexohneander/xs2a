package xs2a

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	authorizationURL = "oauth2/authorize"
)

type (
	AuthorizationResponse struct {
		Location string `json:"location"`
	}
)

// Authorize sends an authorization request to the XS2A API.
// It returns a list of AuthorizationResponse and an error if any.
func (c *Client) Authorize() (resp []AuthorizationResponse, err error) {
	params := map[string]string{
		"client_id":      c.clientId,
		"scope":          "DEDICATED_AISP",
		"code_challenge": c.codeChallenge,
		"redirect_uri":   "https://tpp.com/redirect",
		"response_type":  "CODE",
		"state":          "1fL1nn7m9a",
	}

	res, err := c.do(http.MethodGet, authorizationURL, params)
	if err != nil {
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resp, err
	}

	if res.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("status: %s", res.Status)
	}

	if err = json.Unmarshal(body, &resp); err != nil {
		return resp, err
	}
	return
}
