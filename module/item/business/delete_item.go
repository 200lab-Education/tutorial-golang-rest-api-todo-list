package todobiz

import (
	"context"
	todomodel "first-app/module/item/model"
)

type DeleteTodoItemStorage interface {
	FindItem(
		ctx context.Context,
		condition map[string]interface{},
	) (*todomodel.ToDoItem, error)

	DeleteItem(
		ctx context.Context,
		condition map[string]interface{},
	) error
}

type deleteBiz struct {
	store DeleteTodoItemStorage
}

func NewDeleteToDoItemBiz(store DeleteTodoItemStorage) *deleteBiz {
	return &deleteBiz{store: store}
}

func (biz *deleteBiz) DeleteItem(
	ctx context.Context,
	condition map[string]interface{},
) error {
	// step 1: Find item by conditions
	_, err := biz.store.FindItem(ctx, condition)

	if err != nil {
		return err
	}

	// Step 2: call to storage to delete item
	if err := biz.store.DeleteItem(ctx, condition); err != nil {
		return err
	}

	return nil
}
