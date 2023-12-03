package postgresql

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Querier
}

type SQLRepository struct {
	db *pgxpool.Pool
	*Queries
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &SQLRepository{
		db:      db,
		Queries: New(db),
	}
}
