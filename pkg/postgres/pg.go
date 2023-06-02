package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host     string
	Password string
	Port     string
	Name     string
	Username string
}

func InitPool(config Config) (*pgxpool.Pool, error) {
	baseDsn := "postgres://%s:%s@%s:%s/%s"
	dsn := fmt.Sprintf(baseDsn, config.Username, config.Password, config.Host, config.Port, config.Name)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("error creating pool: %w", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error pinging db: %w", err)
	}

	return pool, nil
}
