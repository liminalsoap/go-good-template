package usecase

import (
	"todo/internal/entitiy"
)

type TaskUseCase struct {
	repo TaskRepo
}

func NewTaskUseCase(r TaskRepo) *TaskUseCase {
	return &TaskUseCase{repo: r}
}

func (uc TaskUseCase) Create(task entitiy.Task) error {
	if err := task.CheckLength(); err != nil {
		return err
	}
	return uc.repo.Create(task)
}

func (uc TaskUseCase) Update(id uint, task entitiy.Task) (entitiy.Task, error) {
	if err := task.CheckLength(); err != nil {
		return entitiy.Task{}, err
	}
	return uc.repo.Update(id, task)
}

func (uc TaskUseCase) Delete(id uint) error {
	return uc.repo.DeleteById(id)
}

func (uc TaskUseCase) GetById(id uint) (entitiy.Task, error) {
	return uc.repo.GetById(id)
}

func (uc TaskUseCase) List() (*[]entitiy.Task, error) {
	return uc.repo.List()
}
