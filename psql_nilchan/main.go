package main

import (
	"context"
	"fmt"
	"psql_stady/feature_postgres/simple_connection"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(conn, "успешный таргет")
	/*
		if err := simple_sql.CreateTable(ctx, conn); err != nil {
			panic(err)
		}
	*/
	/*	if err := simple_sql.InsertRow(ctx,
			conn,
			"Покормить Лютика",
			"Дать ему 30 грамм корма",
			false,
			time.Now(),
		); err != nil {
			panic(err)
		}
	*/

	/*
		if err := simple_sql.DeleteRow(ctx,
			conn,
			4,
		); err != nil {
			panic(err)
		}
	*/

	/*
	   tasks, err := simple_sql.SelectRows(ctx, conn)

	   	if err != nil {
	   		panic(err)
	   	}

	   fmt.Println(tasks)

	   taskExample, err := simple_sql.SelectById(ctx, conn, 1)

	   	if err != nil {
	   		panic(err)
	   	}

	   	if err := simple_sql.UpdateRowTask(ctx, conn, task); err != nil {
	   		panic(err)
	   	}

	   fmt.Println(taskExample)

	   /*

	   		for _, task := range tasks {
	   			if task.ID == 1 {
	   				task.Title = "Покормить Лютика"
	   				task.Description = "Дать ему 30 грамм корма"
	   				task.Completed = true
	   				now := time.Now()
	   				task.CompletedAt = &now

	   				err := simple_sql.UpdateRowTask(ctx, conn, task)
	   				if err != nil {
	   					panic(err)
	   				}
	   			}
	   		}

	   	fmt.Println("succeed!")
	*/
}
