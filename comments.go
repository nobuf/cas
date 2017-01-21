package cas

// Comment object.
type Comment struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	From    User   `json:"from_user"`
	Created int    `json:"created"`
}

// CommentsContainer contains a list of Comments and other information.
type CommentsContainer struct {
	MovieID  MovieID   `json:"movie_id"`
	Count    int       `json:"all_count"`
	Comments []Comment `json:"comments"`
}
