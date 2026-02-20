package main

import (
	"context"
	"fmt"
	"psql_stady/feature_postgres/simple_connection"
	"psql_stady/feature_postgres/simple_sql"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	/*if err := simple_sql.InsertRow(ctx,
		conn,
		"Покормить Лютика",
		"Дать ему 30 грамм корма",
		false,
		time.Now(),
	); err != nil {
		panic(err)
	}
	*/
	if err := simple_sql.UpdateRow(ctx,
		conn,
		4,
	); err != nil {
		panic(err)
	}

	if err := simple_sql.DeleteRow(ctx,
		conn,
		4,
	); err != nil {
		panic(err)
	}

	fmt.Println("succeed!")
}
