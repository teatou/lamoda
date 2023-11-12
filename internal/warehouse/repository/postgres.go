package repository

import (
	"github.com/lib/pq"
	"github.com/teatou/lamoda/internal/warehouse"
	"github.com/teatou/lamoda/pkg/logger"
	storage_postgres "github.com/teatou/lamoda/pkg/storage/postgres"
)

type PostgresRepository struct {
	db     storage_postgres.Postgres
	logger logger.Logger
}

func NewPostgresRepository(db storage_postgres.Postgres, logger logger.Logger) warehouse.Repository {
	return &PostgresRepository{
		db:     db,
		logger: logger,
	}
}

func (p PostgresRepository) ReserveItems(codes []string) error {
	q := `UPDATE item
		SET reserved = true
		WHERE code = ANY($1);
	`
	pgArray := pq.StringArray(codes)
	_, err := p.db.Exec(q, pgArray)
	if err != nil {
		p.logger.Error(err)
		return err
	}
	return nil
}

func (p PostgresRepository) ReleaseItems(codes []string) error {
	q := `UPDATE item
		SET reserved = false
		WHERE code = ANY($1);
	`
	pgArray := pq.StringArray(codes)
	_, err := p.db.Exec(q, pgArray)
	if err != nil {
		p.logger.Error(err)
		return err
	}
	return nil
}

func (p PostgresRepository) GetItemsFromOne(warehouseId int) ([]warehouse.GetItemsFromOneResponse, error) {
	var data []warehouse.GetItemsFromOneResponse
	q := `SELECT item.title as item_title, item.code as item_code, item.reserved as reserved, warehouse_item.quantity as quantity, warehouse.title as warehouse_title
		FROM item
		JOIN warehouse_item ON item.id = warehouse_item.item_id
		JOIN warehouse ON warehouse.id = warehouse_item.warehouse_id
		WHERE warehouse.id = $1;
	`
	err := p.db.Select(&data, q, warehouseId)
	if err != nil {
		p.logger.Error(err)
		return nil, err
	}

	return data, nil
}

func (p PostgresRepository) GetItemRemainder(itemId int) error {
	q := `SELECT item.id, item.title, warehouse.id, warehouse.title, warehouse.available, warehouse_item.quantity
		FROM item
		JOIN warehouse_item ON item.id = warehouse_item.item_id
		JOIN warehouse ON warehouse.id = warehouse_item.warehouse_id
		WHERE item.id = $1
		AND warehouse.available = true;
	`
	_, err := p.db.Exec(q, itemId)
	if err != nil {
		p.logger.Error(err)
		return err
	}

	return nil
}
