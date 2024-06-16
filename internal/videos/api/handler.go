package api

import (
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
	queryText := r.URL.Query().Get("query")
	if queryText == "" {
		errorResponse(400, utils.ErrorResponseStruct{Message: "query is empty"}, w)
		return
	}
	query.Query = queryText
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
