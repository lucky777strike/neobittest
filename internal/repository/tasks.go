package repository

import (
	"camparser"
)

func (r *Repository) CreateTask(task camparser.Task) (int, error) {
	var id int
	query := "INSERT INTO tasks (query,status,count) values ($1, $2, $3) RETURNING id"
	row := r.db.QueryRow(query, task.Query, task.Status, task.Count)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil

}

func (r *Repository) UpdateTask(task camparser.Task) error {
	query := "UPDATE tasks SET status = $1,count = $2  WHERE id = $3;"
	_, err := r.db.Exec(query, task.Status, task.Count, task.Id)
	if err != nil {
		return err
	}
	return nil

}

func (r *Repository) GetAllTasks() ([]camparser.Task, error) {
	var tasks []camparser.Task
	query := "SELECT * FROM tasks"
	err := r.db.Select(&tasks, query)
	return tasks, err
}

func (r *Repository) RemoveAllTasks() error {
	query := "TRUNCATE TABLE tasks;"
	_, err := r.db.Exec(query)
	return err
}
