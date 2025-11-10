package db
import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
    pool, err := pgxpool.New(ctx, dbURL)
    return pool, err
}
