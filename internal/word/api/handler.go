package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"src/domain"
	"src/internal/word/usecase"
	"src/internal/word/utils"
)

type Handlers interface {
	downloadWord(w http.ResponseWriter, r *http.Request)
	getListWord(w http.ResponseWriter, r *http.Request)
}

type wordHandlers struct {
	wordUseCase usecase.UseCase
}

func (h *wordHandlers) downloadWord(w http.ResponseWriter, r *http.Request) {
	log.Println(":Send request to download word.")

	var word domain.Word
	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		errorResponse(400, utils.ErrorResponseStruct{Message: "name required field."}, w)
		return
	}
	word.ID = uuid.New().String()
	err = h.wordUseCase.Create(&word)
	if err != nil {
		errorResponse(400, utils.ErrorResponseStruct{Message: "word already exist."}, w)
		return
	}
	log.Println("success download word")
	wordResponse(word, w, 200)
}

func (h *wordHandlers) getListWord(w http.ResponseWriter, r *http.Request) {
	log.Println(":Send request to autocomplite.")
	var word domain.Word
	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		errorResponse(400, utils.ErrorResponseStruct{"error"}, w)
		return
	}
	list, err := h.wordUseCase.All(&word)
	if err != nil {
		errorResponse(400, utils.ErrorResponseStruct{"error with query"}, w)
		return
	}
	response := map[string][]string{
		"suggestion": list,
	}
	wordResponse(response, w, 200)
}

func NewWordHandlers(metaUC *usecase.UseCase) Handlers {
	return &wordHandlers{wordUseCase: *metaUC}
}
