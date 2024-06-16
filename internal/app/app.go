package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"os"
	"os/signal"
	"src/config"
	api2 "src/internal/videos/api"
	repository2 "src/internal/videos/repository"
	usecase2 "src/internal/videos/usecase"
	"src/internal/word/api"
	"src/internal/word/repository"
	"src/internal/word/usecase"
)

type App struct {
	pgpool *pgxpool.Pool
	conf   *config.AppConfig
}

func NewApp(pgpool *pgxpool.Pool, config *config.AppConfig) *App {
	/* Создание app сервиса */
	return &App{pgpool: pgpool, conf: config}
}

func (a *App) Run() {
	/* Метод запускающий инициализацию репозиториев, юзкейсов, роутов, клиента Minio */

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	wordRepo := repository.NewMetadataRepository(a.pgpool)
	videoRepo := repository2.NewVideoRepository(a.pgpool)
	wordService := usecase.NewWordUseCase(wordRepo, *a.conf)
	videoService := usecase2.NewVideoUseCase(videoRepo, *a.conf)
	metadataHandlers := api.NewWordHandlers(&wordService)
	videoHandlers := api2.NewVideoHandlers(&videoService)

	r := api.Router(metadataHandlers)
	proxyRouter := api2.Router(videoHandlers)
	proxyRouter.Mount("/word", r)
	log.Println(":Route initialization success.")
	go func() {
		err := http.ListenAndServe(a.conf.Server.Port, proxyRouter)
		if err != nil {
			log.Fatal(err)
		}
	}()
	<-ctx.Done()
	log.Println(":Server stopped.")
	cancel()
}
