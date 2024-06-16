package domain

type Word struct {
	ID   string `json:"id,omitempty"`
	Text string `json:"text"`
}

type Words struct {
	Words []string `json:"words"`
}
