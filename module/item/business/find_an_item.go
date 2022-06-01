package todobiz

import (
	"context"
	todomodel "first-app/module/item/model"
)

type FindTodoItemStorage interface {
	FindItem(
		ctx context.Context,
		condition map[string]interface{},
	) (*todomodel.ToDoItem, error)
}

type findBiz struct {
	store FindTodoItemStorage
}

func NewFindToDoItemBiz(store FindTodoItemStorage) *findBiz {
	return &findBiz{store: store}
}

func (biz *findBiz) FindAnItem(ctx context.Context, condition map[string]interface{}) (*todomodel.ToDoItem, error) {
	itemData, err := biz.store.FindItem(ctx, condition)

	if err != nil {
		return nil, err
	}

	return itemData, nil
}
