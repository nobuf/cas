package cas

type Movie struct {
	ID string `json:"id"`
}

type MovieContainer struct {
	Movie       Movie `json:"movie"`
	Broadcaster User  `json:"broadcaster"`
	// TODO tags
}

type Movies struct {
	Movies []MovieContainer `json:"movies"`
}
