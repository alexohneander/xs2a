package xs2a

import (
	"encoding/json"
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

func (c *Client) Authorize() (resp []AuthorizationResponse, err error) {
	params := map[string]string{
		"client_id":      c.clientId,
		"scope":          "DEDICATED_AISP",
		"code_challenge": c.codeChallenge,
		"redirect_uri":   "http://localhost:8080/redirect",
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

	if err = json.Unmarshal(body, &resp); err != nil {
		return resp, err
	}
	return
}
