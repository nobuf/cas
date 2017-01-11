package cas

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var TWITCASTING_API string = "https://apiv2.twitcasting.tv"
var TWITCASTING_API_CLIENT_ID string = os.Getenv("TWITCASTING_API_CLIENT_ID")
var TWITCASTING_API_CLIENT_SECRET string = os.Getenv("TWITCASTING_API_CLIENT_SECRET")

func basicAuthToken() string {
	data := TWITCASTING_API_CLIENT_ID + ":" + TWITCASTING_API_CLIENT_SECRET
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Lives() (Movies, error) {
	// TODO support parameters
	req, _ := http.NewRequest("GET",
		TWITCASTING_API+"/search/lives?type=recommend&lang=ja&limit=50", nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Api-Version", "2.0")
	// TODO support access_token
	req.Header.Add("Authorization", "Basic "+basicAuthToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Movies{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Movies{}, fmt.Errorf("Error: %s", resp.Status)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	movies := Movies{}
	err = json.Unmarshal(body, &movies)
	if err != nil {
		return Movies{}, err
	}
	return movies, nil
}
