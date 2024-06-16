package usecase

import (
	"log"
	"src/config"
	"src/domain"
	repository2 "src/internal/videos/repository"
)

type UseCase interface {
	FilterVectorID(video []string) ([]domain.VideoResponse, error)
}

type videoUseCase struct {
	videoRepo repository2.Repository
	appConfig config.AppConfig
}

func (v videoUseCase) FilterVectorID(video []string) ([]domain.VideoResponse, error) {
	videoMeta, err := v.videoRepo.FilterVectorID(video)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return videoMeta, nil
}

func NewVideoUseCase(videoRepa repository2.Repository, appConfig config.AppConfig) UseCase {
	return &videoUseCase{videoRepo: videoRepa, appConfig: appConfig}
}
