package usecase

import (
	"context"
	"todo/internal/entitiy"
)

type UseCases struct {
	Task Task
}

type Task interface {
	Create(ctx context.Context, task entitiy.Task) error
	Update(ctx context.Context, id uint, task entitiy.Task) (entitiy.Task, error)
	Delete(ctx context.Context, id uint) error

	GetById(ctx context.Context, id uint) (entitiy.Task, error)
	List(ctx context.Context) (*[]entitiy.Task, error)
}

type TaskRepo interface {
	Create(ctx context.Context, task entitiy.Task) error
	Update(ctx context.Context, id uint, task entitiy.Task) (entitiy.Task, error)
	DeleteById(ctx context.Context, id uint) error

	GetById(ctx context.Context, id uint) (entitiy.Task, error)
	List(ctx context.Context) (*[]entitiy.Task, error)
}
