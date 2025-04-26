package handlers

import (
	"fmt"
	"github.com/Monqwa/smartshop/database"
	"github.com/Monqwa/smartshop/entities"
	"github.com/Monqwa/smartshop/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewListRouter(r *gin.Engine) {
	items := r.Group("/list")
	{
		items.POST("/", CreateList)

		items.PUT("/:list_id", UpdateList)

	}
}

func CreateList(c *gin.Context) {
	var list entities.CreateList
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var modelList models.List
	modelList.ID = uuid.New()
	modelList.Title = list.Title
	database.DB.Create(&modelList)
	c.JSON(201, gin.H{"message": "List created", "id": modelList.ID})
}

func UpdateList(c *gin.Context) {
	id := c.Param("list_id")

	var list entities.UpdateList
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var modelList = models.List{}
	database.DB.Where("id = ?", id).First(&modelList)
	fmt.Println(modelList, id)
	if list.Title != "" {
		modelList.Title = list.Title
		database.DB.Save(&modelList)
	}

	c.JSON(200, gin.H{"message": "List updated"})
}
