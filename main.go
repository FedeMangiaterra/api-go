package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

var users = []User{
	{ID: "1", Name: "Lionel Messi", Email: "messi@hotmail.com", Age: 37},
	{ID: "2", Name: "Egg Acuna", Email: "huevo@gmail.com", Age: 33},
	{ID: "3", Name: "Father Gomez", Email: "papu@gmail.com", Age: 36},
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	for _, user := range users {
		if user.ID == id {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func CreateUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		return
	}

	if user.ID == "" || user.Name == "" || user.Email == "" || user.Age == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing required fields"})
		return
	}

	users = append(users, user)
	c.IndentedJSON(http.StatusCreated, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, nil)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	port := os.Getenv("PORT")

	router.GET("/users", GetUsers)
	router.GET("/users/:id", GetUser)

	router.POST("/users", CreateUser)

	router.DELETE("/users/:id", DeleteUser)

	err = router.Run(port)
	if err != nil {
		panic("Failed to start server on port " + port + " due to error: " + err.Error())
	}
}
