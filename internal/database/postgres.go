package postgres

import (
	"context"

	"github.com/DLzer/go-echo-boilerplate/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

// const (
// 	maxOpenConns    = 60
// 	connMaxLifetime = 120
// 	maxIdleConns    = 30
// 	connMaxIdleTime = 20
// )

// Return new Postgresql db instance
func NewPsqlDB(c *config.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), c.Postgres.PostgresqlDSN)
	if err != nil {
		return nil, err
	}

	pool.Config().AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	return pool, nil
}
