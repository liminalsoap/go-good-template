package entitiy

import (
	"errors"
	"time"
)

type Task struct {
	Id          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

const maxLength = 5

func (t Task) CheckLength() error {
	if len(t.Title) < maxLength {
		return errors.New("length of title must be > 5")
	}
	return nil
}
