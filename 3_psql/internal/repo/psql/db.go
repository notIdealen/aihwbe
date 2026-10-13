package psql

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	*pgxpool.Pool
	schema          string
	table           string
	timerForContext time.Duration
}

func NewPostgresPool(ctx context.Context, config *PostgresConfig) *PostgresDB {
	url := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", config.User, config.Password, config.DBName)
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		// panic("Bad new pool with postgress")
	}

	tctx, cf := context.WithTimeout(ctx, time.Duration(config.Timeout))
	defer cf()
	if err := pool.Ping(tctx); err != nil {
		// panic("Bad PING with postgress")
	}

	return &PostgresDB{Pool: pool,
		table:           pgx.Identifier{config.Schema, config.Table}.Sanitize(),
		timerForContext: time.Second * time.Duration(config.Timeout),
	}
}
