package cas

type MovieID string

type Movie struct {
	ID               MovieID `json:"id"`
	UserID           string  `json:"user_id"`
	Title            string  `json:"title"`
	Subtitle         string  `json:"subtitle"`
	Link             string  `json:"link"`
	Live             bool    `json:"is_live"`
	Recorded         bool    `json:"is_recorded"`
	CommentCount     int     `json:"comment_count"`
	LargeThumbnail   string  `json:"large_thumbnail"`
	SmallThumbnail   string  `json:"small_thumbnail"`
	Country          string  `json:"country"`
	Duration         int     `json:"duration"`
	Created          int     `json:"created"`
	Collabo          bool    `json:"is_collabo"`
	Protected        bool    `json:"is_protected"`
	MaxViewCount     int     `json:"max_view_count"`
	CurrentViewCount int     `json:"current_view_count"`
	TotalViewCount   int     `json:"total_view_count"`
	HLS              string  `json:"hls_url"`
}

type MovieContainer struct {
	Movie       Movie    `json:"movie"`
	Broadcaster User     `json:"broadcaster"`
	Tags        []string `json:"tags"`
}

type Movies struct {
	Movies []MovieContainer `json:"movies"`
}
