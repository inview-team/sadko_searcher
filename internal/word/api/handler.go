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

	var words domain.Words
	err := json.NewDecoder(r.Body).Decode(&words)
	if err != nil {
		errorResponse(400, utils.ErrorResponseStruct{Message: "name required field."}, w)
		return
	}
	for _, attr := range words.Words {
		var word domain.Word
		word.ID = uuid.New().String()
		word.Text = attr
		err = h.wordUseCase.Create(&word)
		if err != nil {
			errorResponse(400, utils.ErrorResponseStruct{Message: "word already exist."}, w)
			return
		}
	}
	log.Println("success download word")
	wordResponse(words, w, 200)
}

func (h *wordHandlers) getListWord(w http.ResponseWriter, r *http.Request) {
	log.Println(":Send request to autocomplite.")
	var word domain.Word
	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		errorResponse(400, utils.ErrorResponseStruct{Message: "error with decode body."}, w)
		return
	}
	list, err := h.wordUseCase.All(&word)
	if err != nil {
		errorResponse(400, utils.ErrorResponseStruct{Message: "error with list query."}, w)
		return
	}
	if len(list) == 0 {
		word.ID = uuid.New().String()
		err = h.wordUseCase.Create(&word)
		if err != nil {
			log.Println(err)
			errorResponse(400, utils.ErrorResponseStruct{Message: "error with save word in suggestions"}, w)
		}
		list = append(list, word.Text)
	}
	response := map[string][]string{
		"suggestion": list,
	}
	wordResponse(response, w, 200)
}

func NewWordHandlers(metaUC *usecase.UseCase) Handlers {
	return &wordHandlers{wordUseCase: *metaUC}
}
