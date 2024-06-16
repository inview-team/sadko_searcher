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
	//var videos []string
	//err := json.NewDecoder(r.Body).Decode(&videos)
	//if err != nil {
	//	errorResponse(400, utils.ErrorResponseStruct{Message: "error"}, w)
	//	return
	//}
	videos := []string{"420c0854-ced1-497b-894f-a87c63ba7f64", "620c0854-ced1-497b-894f-a87c63ba7f64"}
	list, err := h.videoUseCase.FilterVectorID(videos)
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
