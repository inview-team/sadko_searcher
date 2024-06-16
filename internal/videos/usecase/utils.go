package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"src/domain"
)

type Payload struct {
	SearchText string `json:"search_text"`
}

func VideoProcessorRequest(query string, endpoint string) ([]string, error) {
	body := []byte(query)
	searchText := Payload{SearchText: query}
	jsonPayload, err := json.Marshal(searchText)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil, err
	}
	req, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer req.Body.Close()
	body, err = io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var videoProcessorResponse domain.VideoSearch
	err = json.Unmarshal(body, &videoProcessorResponse)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return videoProcessorResponse.VectorIDs, nil

}
