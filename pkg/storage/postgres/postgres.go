package storage_postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/teatou/lamoda/internal/config"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

const (
	_defaultConnectionTimeout = time.Second
	_maxConnections           = int32(10)
	_minConnections           = int32(1)
	_maxConnectionLifeTime    = time.Second * 300
	_maxIdleLifeTime          = time.Second * 15
)

type Postgres interface {
	Stats() *pgxpool.Stat
	Query(query string, args ...any) (pgx.Rows, error)
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...any) (pgconn.CommandTag, error)
	QueryRow(query string, args ...interface{}) pgx.Row
	TxRunner
}

type Pool struct {
	db *pgxpool.Pool
}

func InitPsqlDB(c config.PostgresConfig) (Postgres, error) {
	connectionUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DbName)
	connectionUrl += fmt.Sprintf(" pool_max_conns=%d pool_min_conns=%d pool_max_conn_lifetime=%v pool_max_conn_idle_time=%v",
		_maxConnections, _minConnections, _maxConnectionLifeTime, _maxIdleLifeTime)

	result, err := pgxpool.Connect(context.Background(), connectionUrl)
	if result == nil {
		log.Printf("POSTGRES CONNECTION(%s) ERROR: %s\n", connectionUrl, err.Error())
		return nil, err
	}

	return &Pool{db: result}, nil
}

func (p Pool) Stats() *pgxpool.Stat {
	return p.db.Stat()
}

func (p Pool) Begin(ctx context.Context) (pgx.Tx, error) {
	return p.db.Begin(ctx)
}

func (p Pool) Query(query string, args ...any) (pgx.Rows, error) {
	return p.db.Query(context.Background(), query, args[:]...)
}

func (p Pool) Select(dest interface{}, query string, args ...interface{}) error {
	rows, err := p.db.Query(context.Background(), query, args[:]...)
	if err != nil {
		return err
	}
	return pgxscan.DefaultAPI.ScanAll(dest, rows)
}

func (p Pool) Exec(query string, args ...any) (pgconn.CommandTag, error) {
	return p.db.Exec(context.Background(), query, args[:]...)
}

func (p Pool) QueryRow(query string, args ...interface{}) pgx.Row {
	return p.db.QueryRow(context.Background(), query, args[:]...)
}
