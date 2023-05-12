package repository

import (
	"context"
	"errors"
	"todo/internal/entitiy"
	"todo/pkg/postgres"
)

type TaskRepo struct {
	*postgres.Postgres
}

func NewTaskRepo(pg *postgres.Postgres) *TaskRepo {
	return &TaskRepo{pg}
}

func (t TaskRepo) Create(task entitiy.Task) error {
	_, err := t.Conn.Exec(
		context.Background(),
		"INSERT INTO tasks(title, description) values ($1, $2)",
		task.Title,
		task.Description,
	)
	return err
}

func (t TaskRepo) Update(id uint, task entitiy.Task) (entitiy.Task, error) {
	_, err := t.Conn.Exec(
		context.Background(),
		"UPDATE tasks SET updated_at=now(), title=$1, description=$2 WHERE id = $3",
		task.Title,
		task.Description,
		id,
	)
	if err != nil {
		return entitiy.Task{}, err
	}

	updatedTask, err := t.GetById(id)
	if err != nil {
		return entitiy.Task{}, err
	}
	return updatedTask, nil
}

func (t TaskRepo) DeleteById(id uint) error {
	flag, err := t.Conn.Exec(
		context.Background(),
		"DELETE FROM tasks WHERE id = $1",
		id,
	)
	if flag.RowsAffected() == 0 {
		return errors.New("not found task")
	}
	return err
}

func (t TaskRepo) GetById(id uint) (entitiy.Task, error) {
	var task entitiy.Task
	err := t.Conn.QueryRow(
		context.Background(),
		"SElECT * FROM tasks WHERE id = $1", id).Scan(
		&task.Id,
		&task.CreatedAt,
		&task.UpdatedAt,
		&task.Title,
		&task.Description,
	)
	if err != nil {
		return entitiy.Task{}, err
	}
	return task, nil
}

func (t TaskRepo) List() (*[]entitiy.Task, error) {
	rows, err := t.Conn.Query(context.Background(), "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	var tasks []entitiy.Task
	for rows.Next() {
		var task entitiy.Task
		err := rows.Scan(
			&task.Id,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.Title,
			&task.Description,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &tasks, nil
}
