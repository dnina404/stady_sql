package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context,
	conn *pgx.Conn,
	id int,
) error {
	sqlQuery := `
	Delete from tasks
	WHERE id = $1
	`

	_, err := conn.Exec(ctx, sqlQuery, id)

	return err
}
