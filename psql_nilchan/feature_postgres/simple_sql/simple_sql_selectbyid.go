package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func SelectById(ctx context.Context,
	conn *pgx.Conn,
	id int,
) (TaskModel, error) {
	sqlQuery := `
	SELECT * FROM tasks
	WHERE id = $1
	`

	rows, err := conn.Query(ctx, sqlQuery, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var task TaskModel
	for rows.Next() {
		var task TaskModel

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.CompletedAt,
		)
		if err != nil {
			return TaskModel{}, err
		}
	}

	return task, nil
}
