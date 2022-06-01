package todostorage

import (
	"context"
	todomodel "first-app/module/item/model"
)

func (s *mysqlStorage) UpdateItem(
	ctx context.Context,
	condition map[string]interface{},
	dataUpdate *todomodel.ToDoItem,
) error {
	if err := s.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}
