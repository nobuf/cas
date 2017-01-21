package cas

// SearchRecommendLives retrieves recommended Moves on live
// See http://apiv2-doc.twitcasting.tv/#search-live-movies
func (api *Client) SearchRecommendLives() (Movies, error) {
	movies := Movies{}
	// TODO support parameters
	err := get(api, "/search/lives?type=recommend&lang=ja&limit=50", &movies)
	if err != nil {
		return Movies{}, err
	}
	return movies, nil
}
