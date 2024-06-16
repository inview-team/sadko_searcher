package api

import (
	"encoding/json"
	"log"
	"net/http"
	"src/domain"
	"src/internal/videos/usecase"
	"src/internal/word/utils"
)

type Handlers interface {
	filterVectorID(w http.ResponseWriter, r *http.Request)
}

type videoHandlers struct {
	videoUseCase usecase.UseCase
}

func (h *videoHandlers) filterVectorID(w http.ResponseWriter, r *http.Request) {
	log.Println(":Send request to filter vector id.")
	var query domain.QuerySearch
	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		errorResponse(400, utils.ErrorResponseStruct{Message: "error with decode query."}, w)
		return
	}
	list, err := h.videoUseCase.FilterVectorID(query.Query)
	if err != nil {
		errorResponse(400, utils.ErrorResponseStruct{Message: "error"}, w)
		return
	}
	response := map[string][]domain.VideoResponse{
		"result": list,
	}
	videoResponse(response, w, 200)
}

func NewVideoHandlers(metaUC *usecase.UseCase) Handlers {
	return &videoHandlers{videoUseCase: *metaUC}
}
