package models

import (
	"fmt"
	"todo-api/db"
)

func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()
	//fmt.Println(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}
