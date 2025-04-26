package handlers

import (
	"github.com/Monqwa/smartshop/database"
	"github.com/Monqwa/smartshop/entities"
	"github.com/Monqwa/smartshop/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewItemRouter(r *gin.Engine) {
	items := r.Group("/items")
	{
		items.POST("/:list_id", CreateItem)

		items.GET("/:list_id", GetItems)

		items.PUT("/:id", UpdateItem)

		items.DELETE("/:id", DeleteItem)

	}
}

func CreateItem(c *gin.Context) {
	listId := c.Param("list_id")

	var item entities.CreateItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var modelItem models.Item
	listIdUuid, _ := uuid.Parse(listId)

	modelItem.ID = uuid.New()
	modelItem.Name = item.Name
	modelItem.ListID = listIdUuid
	modelItem.IsBuy = false

	database.DB.Create(&modelItem)

	c.JSON(201, gin.H{"message": "Item created"})
}

func GetItems(c *gin.Context) {
	listId := c.Param("list_id")
	var items []models.Item
	database.DB.Where("list_id = ?", listId).Find(&items)
	c.JSON(200, items)
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")

	var item entities.UpdateItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var modelItem = models.Item{}
	database.DB.Where("id = ?", id).First(&modelItem)
	if item.Name != "" {
		modelItem.Name = item.Name
	}
	if item.IsBuy != nil {
		modelItem.IsBuy = *item.IsBuy
	}
	database.DB.Save(&modelItem)
	c.JSON(200, gin.H{"message": "Item updated", "id": id})

}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	var modelItem = models.Item{}
	database.DB.Where("id = ?", id).First(&modelItem)
	database.DB.Delete(&modelItem)
	c.JSON(200, gin.H{"message": "Item deleted", "id": id})
}
