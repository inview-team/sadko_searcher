package usecase

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"src/domain"
)

type SearchRequest struct {
	Query string `json:"query"`
}

func VideoProcessorRequest(query string, endpoint string) ([]string, error) {
	searchText := SearchRequest{Query: query}
	jsonPayload, err := json.Marshal(searchText)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return nil, err
	}
	req, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var videoProcessorResponse domain.VideoSearch
	err = json.Unmarshal(body, &videoProcessorResponse)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return videoProcessorResponse.VectorIDs, nil
}
