package pg

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	db *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
	errorDb    error
)

func NewPG(ctx context.Context, connString string) (*Postgres, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			log.Fatal("unable to create connection pool: %w", err)
			errorDb = err
			return
		}

		pgInstance = &Postgres{db}
	})

	if errorDb != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", errorDb)
	}

	return pgInstance, nil
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.db.Close()
}

func (pg *Postgres) BeginTx(ctx context.Context) (transaction pgx.Tx, err error) {
	return pg.db.BeginTx(ctx, pgx.TxOptions{})
}
