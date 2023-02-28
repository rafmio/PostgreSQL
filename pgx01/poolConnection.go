package pg

import (
  "context"
  "fmt"
  "sync"

  "github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
  db *pbxpool.Pool
}

var (
  pgInstance *postgres
  pgOnce sync.Once
)

func NewPG(ctx context.Context, connString string) (*postgres, error) {
  pgOnce.Do(func() {
    db, err := pgxpool.New(ctx, connString)
    if err != nil {
      return fmt.Errorf("unable to create connection pool: %w", err)
    }

    pgInstance = &postgres{db}
  })

  return pgInstance, nil
}

func (pg *postgres) Ping(ctx context.Context) error {
  return pg.db.Ping(ctx)
}

func (pg *postgres) Close() {
  pg.db.Close()
}
