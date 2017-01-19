package cas

var TWITCASTING_API string = "https://apiv2.twitcasting.tv"

type Client struct {
	clientID     string
	clientSecret string
	endpoint     string
}

// TODO support OAuth token
func New(clientID string, clientSecret string) *Client {
	c := &Client{
		clientID:     clientID,
		clientSecret: clientSecret,
		endpoint:     TWITCASTING_API,
	}
	return c
}

// Override API host and the root path.
func (api *Client) SetEndpoint(endpoint string) {
	api.endpoint = endpoint
}
