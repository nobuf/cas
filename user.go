package cas

type User struct {
	ID          string  `json:"id"`
	ScreenID    string  `json:"screen_id"`
	Name        string  `json:"name"`
	Image       string  `json:"image"`
	Profile     string  `json:"profile"`
	LastMovieID MovieID `json:"last_movie_id"`
	Live        bool    `json:"is_live"`
	Created     int     `json:"created"`
}
