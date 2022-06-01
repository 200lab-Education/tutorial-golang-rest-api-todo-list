package todobiz

import (
	"context"
	todomodel "first-app/module/item/model"
)

type ListTodoItemStorage interface {
	ListItem(
		ctx context.Context,
		condition map[string]interface{},
		paging *todomodel.DataPaging,
	) ([]todomodel.ToDoItem, error)
}

type listBiz struct {
	store ListTodoItemStorage
}

func NewListToDoItemBiz(store ListTodoItemStorage) *listBiz {
	return &listBiz{store: store}
}

func (biz *listBiz) ListItems(ctx context.Context,
	condition map[string]interface{},
	paging *todomodel.DataPaging,
) ([]todomodel.ToDoItem, error) {
	result, err := biz.store.ListItem(ctx, condition, paging)

	if err != nil {
		return nil, err
	}

	return result, err
}
