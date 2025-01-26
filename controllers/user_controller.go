package controllers

import (
	"net/http"
	"strconv"

	"api-test/models"
	"api-test/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := services.GetUsers()
	c.IndentedJSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	user, found := services.GetUserByID(id)
	if found {
		c.IndentedJSON(http.StatusOK, user)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Name == "" || user.Email == "" || user.Age == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing required fields"})
		return
	}

	user = services.AddUser(user)
	c.IndentedJSON(http.StatusCreated, user)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	found := services.DeleteUser(id)
	if found {
		c.IndentedJSON(http.StatusNoContent, nil)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedUser models.User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updatedUser.Name == "" && updatedUser.Email == "" && updatedUser.Age == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing required fields"})
		return
	}

	updatedUser, found := services.UpdateUser(id, updatedUser)
	if found {
		c.IndentedJSON(http.StatusOK, updatedUser)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
}
