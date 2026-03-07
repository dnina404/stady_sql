package simple_sql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func CompleteTask(ctx context.Context,
	conn *pgx.Conn,
	id int,
	boolean bool,
	time time.Time,
) (TaskModel, error) {
	sqlQuery := `
	UPDATE tasks
	SET completed = $2,
    completed_at = $3
	WHERE id = $1;
	`

	rows, err := conn.Query(ctx, sqlQuery, id, boolean, time)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	task, err := SelectById(ctx, conn, id)
	if err != nil {
		panic(err)
	}
	return task, nil
}
