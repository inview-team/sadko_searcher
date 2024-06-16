package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
	"log"
	"src/domain"
)

type Repository interface {
	FilterVectorID(video []string) ([]domain.VideoResponse, error)
}

type videoRepository struct {
	db *pgxpool.Pool
}

func (v videoRepository) FilterVectorID(video []string) ([]domain.VideoResponse, error) {
	var videoMeta []domain.VideoResponse
	rows, err := v.db.Query(context.Background(), selectArray, pq.Array(video))
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var videoSchema domain.VideoResponse
		if err := rows.Scan(&videoSchema.Description, &videoSchema.Url); err != nil {
			log.Fatal(err)
		}
		videoMeta = append(videoMeta, videoSchema)
	}
	return videoMeta, nil
}

func NewVideoRepository(pgpool *pgxpool.Pool) Repository {
	return &videoRepository{db: pgpool}
}
