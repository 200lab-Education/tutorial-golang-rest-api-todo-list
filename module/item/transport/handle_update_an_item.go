package todotrpt

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	todobiz "first-app/module/item/business"
	todomodel "first-app/module/item/model"
	todostorage "first-app/module/item/storage"
)

func HandleUpdateAnItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dataItem todomodel.ToDoItem

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := todostorage.NewMySQLStorage(db)
		biz := todobiz.NewUpdateToDoItemBiz(storage)

		if err := biz.UpdateItem(c.Request.Context(), map[string]interface{}{"id": id}, &dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
