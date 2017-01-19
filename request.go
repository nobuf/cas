package cas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func get(api *Client, path string, responseFormat interface{}) error {
	// TODO support parameters
	req, _ := http.NewRequest("GET",
		api.endpoint+path, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Api-Version", "2.0")
	req.Header.Add("Authorization", api.auth())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("Error: %s", resp.Status)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &responseFormat)
	if err != nil {
		return err
	}
	return nil
}
