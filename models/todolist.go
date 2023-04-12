package models

import (
	"jubo-hw/schemas"
)

func GetAllTODOItems() ([]schemas.TodoItem, error) {
	rows, err := db.Query("SELECT * FROM todo_list")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todoItems := []schemas.TodoItem{}
	for rows.Next() {
		var todoItem schemas.TodoItem
		err := rows.Scan(&todoItem.ID, &todoItem.Title, &todoItem.Description, &todoItem.Completed, &todoItem.CreatedAt)
		if err != nil {
			return nil, err
		}
		todoItems = append(todoItems, todoItem)
	}
	return todoItems, nil
}

func GetTODOItem(id int) (schemas.TodoItem, error) {
	var todoItem schemas.TodoItem
	err := db.QueryRow("SELECT * FROM todo_list WHERE id = ?", id).Scan(&todoItem.ID, &todoItem.Title, &todoItem.Description, &todoItem.Completed, &todoItem.CreatedAt)
	if err != nil {
		return schemas.TodoItem{}, err
	}
	return todoItem, nil
}

func CreateTODOItem(todoItem schemas.CreateTodoItemBody) error {
	_, err := db.Exec("INSERT INTO todo_list (title, description, completed) VALUES (?, ?, ?)", todoItem.Title, todoItem.Description, todoItem.Completed)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTODOItem(id int) error {
	if _, err := GetTODOItem(id); err != nil {
		return err
	}
	if _, err := db.Exec("UPDATE todo_list SET completed = ? WHERE id = ?", true, id); err != nil {
		return err
	}
	return nil
}

func DeleteTODOItem(id int) error {
	if _, err := GetTODOItem(id); err != nil {
		return err
	}
	if _, err := db.Exec("DELETE FROM todo_list WHERE id = ?", id); err != nil {
		return err
	}
	return nil
}
