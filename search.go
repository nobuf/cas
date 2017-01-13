package cas

func (api *Client) SearchRecommendLives() (Movies, error) {
	movies := Movies{}
	err := get("/search/lives?type=recommend&lang=ja&limit=50", api.auth(), &movies)
	if err != nil {
		return Movies{}, err
	}
	return movies, nil
}
