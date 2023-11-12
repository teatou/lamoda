package repository

import (
	"github.com/teatou/lamoda/internal/warehouse"
	"github.com/teatou/lamoda/pkg/logger"
	storage_postgres "github.com/teatou/lamoda/pkg/storage/postgres"
)

type PostgresRepository struct {
	db     storage_postgres.Postgres
	logger *logger.Logger
}

func NewPostgresRepository(db storage_postgres.Postgres, logger *logger.Logger) warehouse.Repository {
	return &PostgresRepository{
		db:     db,
		logger: logger,
	}
}
