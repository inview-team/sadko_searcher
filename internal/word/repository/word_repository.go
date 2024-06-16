package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"src/domain"
	"strings"
)

type Repository interface {
	Create(word *domain.Word) error
	All(word string) ([]string, error)
}

type wordRepository struct {
	db *pgxpool.Pool
}

func NewMetadataRepository(pgpool *pgxpool.Pool) Repository {
	return &wordRepository{db: pgpool}
}

func (w wordRepository) Create(word *domain.Word) error {
	_, err := w.db.Query(context.Background(), createWord, word.ID, strings.ToLower(word.Text))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (w wordRepository) All(word string) ([]string, error) {
	var wordSlice []string
	rows, err := w.db.Query(context.Background(), suggestionsWord, word+"%")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		var wordSchema domain.Word
		if err := rows.Scan(&wordSchema.ID, &wordSchema.Text); err != nil {
			log.Println(err)
			return nil, err
		}
		wordSlice = append(wordSlice, wordSchema.Text)
	}
	return wordSlice, nil
}
