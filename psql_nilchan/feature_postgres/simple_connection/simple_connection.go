package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// "postgres://YourName:YourPassword@YourHostName:5432/YourDataBaseName"

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres")
}
