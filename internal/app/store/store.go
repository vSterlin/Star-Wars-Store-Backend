package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	db                *sql.DB
	productRepository *ProductRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Product() *ProductRepository {
	if s.productRepository != nil {
		return s.productRepository
	}

	s.productRepository = &ProductRepository{
		store: s,
	}

	return s.productRepository
}
