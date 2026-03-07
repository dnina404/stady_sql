package simple_connection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

// "postgres://YourName:YourPassword@YourHostName:5432/YourDataBaseName"

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	connString := os.Getenv("CONN_STRING")
	return pgx.Connect(ctx, connString)
}
