package usecase

import "todo/internal/entitiy"

type UseCases struct {
	Task Task
}

type Task interface {
	Create(task entitiy.Task) error
	Update(id uint, task entitiy.Task) (entitiy.Task, error)
	Delete(id uint) error

	GetById(id uint) (entitiy.Task, error)
	List() (*[]entitiy.Task, error)
}

type TaskRepo interface {
	Create(task entitiy.Task) error
	Update(id uint, task entitiy.Task) (entitiy.Task, error)
	DeleteById(id uint) error

	GetById(id uint) (entitiy.Task, error)
	List() (*[]entitiy.Task, error)
}
