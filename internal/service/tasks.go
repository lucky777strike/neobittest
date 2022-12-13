package service

import (
	"camparser"
	"fmt"
)

func (s *Service) CreateTask(task camparser.Task) (int, error) {
	return s.repo.CreateTask(task)
}

func (s *Service) UpdateTask(id int, count int, err_count int, errorflag bool) error {
	var status string
	if !errorflag {
		status = fmt.Sprintf("Inserted with %d errors", err_count)
	} else {
		status = "error"
	}
	return s.repo.UpdateTask(camparser.Task{Id: id,
		Status: status,
		Count:  count})
}

func (s *Service) GetAllTasks() ([]camparser.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *Service) RemoveAllTasks() error {
	return s.repo.RemoveAllTasks()
}
