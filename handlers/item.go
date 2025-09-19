package handlers

import (
	"net/http"
	"strconv"

	"takehome-cbi/golang-be/models"

	"github.com/gin-gonic/gin"
)

var items = []models.Item{
	{ID: 1, Name: "Item 1", Price: 100},
	{ID: 2, Name: "Item 2", Price: 200},
}
var nextID = 3

func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

func CreateItem(c *gin.Context) {
	var newItem models.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	newItem.ID = nextID
	nextID++
	items = append(items, newItem)
	c.JSON(http.StatusCreated, newItem)
}

func UpdateItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedItem models.Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, item := range items {
		if item.ID == id {
			items[i].Name = updatedItem.Name
			items[i].Price = updatedItem.Price
			c.JSON(http.StatusOK, items[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
}

func DeleteItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
}
