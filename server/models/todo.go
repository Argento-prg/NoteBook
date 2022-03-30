package models

import "errors"

// структура списков
type List struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

// структура задач
type Todo struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Importance  bool   `json:"importance" db:"importance"`
	Done        bool   `json:"done" db:"done"`
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type UpdateTodoInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Importance  *bool   `json:"importance"`
	Done        *bool   `json:"done"`
}

func (i UpdateTodoInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Importance == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
