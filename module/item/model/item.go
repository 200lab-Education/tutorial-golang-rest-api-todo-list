package todomodel

import (
	"errors"
	"time"
)

var (
	ErrTitleCannotBeBlank       = errors.New("title can not be blank")
	ErrItemNotFound             = errors.New("item not found")
	ErrCannotUpdateFinishedItem = errors.New("can not update finished item")
)

type ToDoItem struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Title     string     `json:"title" gorm:"column:title;"`
	Status    string     `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (ToDoItem) TableName() string { return "todo_items" }

func (item ToDoItem) Validate() error {
	if item.Title == "" {
		return ErrTitleCannotBeBlank
	}

	return nil
}

type DataPaging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *DataPaging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 10
	}
}
