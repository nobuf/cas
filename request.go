package cas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get open custom get request to let user request the api we didn't provide
func (api *Client) Get(path string, responseFormat interface{}) error {
	return get(api, path, responseFormat)
}

// Make an http request with GET method to the API server.
func get(api *Client, path string, responseFormat interface{}) error {
	req, _ := http.NewRequest("GET",
		api.endpoint+path, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Api-Version", "2.0")
	if api.hasAuthToken() {
		req.Header.Add("Authorization", api.auth())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		var requestError *RequestError
		body, _ := ioutil.ReadAll(resp.Body)
		// If it cannot parse to TwitCasting response error, just throw the error status
		if err = json.Unmarshal(body, &requestError); err != nil {
			return fmt.Errorf("http error: %s", resp.Status)
		}
		requestError.StatusCode = resp.StatusCode // assign real http status code
		return requestError
	}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &responseFormat)
	if err != nil {
		return err
	}
	return nil
}
