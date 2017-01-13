package cas

import (
	"encoding/base64"
)

func (api *Client) basicAuthToken() string {
	data := api.clientID + ":" + api.clientSecret
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func (api *Client) auth() string {
	// TODO use access_token if it's available
	return "Basic " + api.basicAuthToken()
}
