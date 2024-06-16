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

type VideoSearch struct {
	VectorIDs []string `json:"vector_ids"`
}

type QuerySearch struct {
	Query string `json:"query"`
}
