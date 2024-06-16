package usecase

import (
	"log"
	"src/config"
	"src/domain"
	"src/internal/word/repository"
)

type UseCase interface {
	Create(word *domain.Word) error
	All(word *domain.Word) ([]string, error)
}

type wordUseCase struct {
	wordRepo  repository.Repository
	appConfig config.AppConfig
}

func NewWordUseCase(metaRepo repository.Repository, appConfig config.AppConfig) UseCase {
	return &wordUseCase{wordRepo: metaRepo, appConfig: appConfig}
}

func (w wordUseCase) Create(word *domain.Word) error {
	err := w.wordRepo.Create(word)
	if err != nil {
		return err
	}
	return nil
}

func (w wordUseCase) All(word *domain.Word) ([]string, error) {
	wordSlice, err := w.wordRepo.All(word.Text)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return wordSlice, nil
}
