package cas

// TwitCastingAPI is the endpoint of API. It's possible to override it, or use SetEndpoint().
var TwitCastingAPI = "https://apiv2.twitcasting.tv"

// Client is the main struct to access all the methods that this package provides.
type Client struct {
	clientID     string
	clientSecret string
	endpoint     string
}

// New initializes Client.
// TODO support OAuth token
func New(clientID string, clientSecret string) *Client {
	c := &Client{
		clientID:     clientID,
		clientSecret: clientSecret,
		endpoint:     TwitCastingAPI,
	}
	return c
}

// SetEndpoint overrides the API host and the root path.
func (api *Client) SetEndpoint(endpoint string) {
	api.endpoint = endpoint
}
