package cas

import (
	"encoding/base64"
)

func (api *Client) basicAuthToken() string {
	data := api.clientID + ":" + api.clientSecret
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func (api *Client) hasAuthToken() bool {
	if len(api.clientID) > 0 && len(api.clientSecret) > 0 {
		return true
	}
	return false
}

// Return Authorization http header
func (api *Client) auth() string {
	// TODO use access_token if it's available
	return "Basic " + api.basicAuthToken()
}
