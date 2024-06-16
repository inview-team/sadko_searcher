package usecase

import (
	"log"
	"src/config"
	"src/domain"
	repository2 "src/internal/videos/repository"
)

type UseCase interface {
	FilterVectorID(query string) ([]domain.VideoResponse, error)
}

type videoUseCase struct {
	videoRepo repository2.Repository
	appConfig config.AppConfig
}

func (v videoUseCase) FilterVectorID(query string) ([]domain.VideoResponse, error) {
	vectorIDs, err := VideoProcessorRequest(query, v.appConfig.Server.EndpointVideoProcessor)
	if err != nil {
		log.Println(nil, err)
		return nil, err
	}
	videoMeta, err := v.videoRepo.FilterVectorID(vectorIDs)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return videoMeta, nil
}

func NewVideoUseCase(videoRepa repository2.Repository, appConfig config.AppConfig) UseCase {
	return &videoUseCase{videoRepo: videoRepa, appConfig: appConfig}
}
