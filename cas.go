package cas

var TWITCASTING_API string = "https://apiv2.twitcasting.tv"

type Client struct {
	clientID string
	clientSecret string
}

// TODO support OAuth token
func New(clientID string, clientSecret string) *Client {
	c := &Client{
		clientID:clientID,
		clientSecret:clientSecret,
	}
	return c
}
