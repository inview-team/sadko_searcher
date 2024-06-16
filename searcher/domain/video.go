package domain

type Video struct {
	ID               string
	Url              string
	Description      string
	RelatedVectorIDs []string
}

type VideoResponse struct {
	Url         string
	Description string
}
