package usecase

import (
	"context"
	"todo/internal/entitiy"
)

type TaskUseCase struct {
	repo TaskRepo
}

func NewTaskUseCase(r TaskRepo) *TaskUseCase {
	return &TaskUseCase{repo: r}
}

func (uc TaskUseCase) Create(ctx context.Context, task entitiy.Task) error {
	if err := task.CheckLength(); err != nil {
		return err
	}
	return uc.repo.Create(ctx, task)
}

func (uc TaskUseCase) Update(ctx context.Context, id uint, task entitiy.Task) (entitiy.Task, error) {
	if err := task.CheckLength(); err != nil {
		return entitiy.Task{}, err
	}
	return uc.repo.Update(ctx, id, task)
}

func (uc TaskUseCase) Delete(ctx context.Context, id uint) error {
	return uc.repo.DeleteById(ctx, id)
}

func (uc TaskUseCase) GetById(ctx context.Context, id uint) (entitiy.Task, error) {
	return uc.repo.GetById(ctx, id)
}

func (uc TaskUseCase) List(ctx context.Context) (*[]entitiy.Task, error) {
	return uc.repo.List(ctx)
}
