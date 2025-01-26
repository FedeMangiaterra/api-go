package services

import "api-test/models"

var users = []models.User{}
var nextID = 1

func GetUsers() []models.User {
	return users
}

func GetUserByID(id int) (models.User, bool) {
	for _, user := range users {
		if user.ID == id {
			return user, true
		}
	}
	return models.User{}, false
}

func AddUser(user models.User) models.User {
	user.ID = nextID
	nextID++
	users = append(users, user)
	return user
}

func DeleteUser(id int) bool {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}

func UpdateUser(id int, updatedUser models.User) (models.User, bool) {
	for i, user := range users {
		if user.ID == id {
			if updatedUser.Name != "" {
				users[i].Name = updatedUser.Name
			}
			if updatedUser.Email != "" {
				users[i].Email = updatedUser.Email
			}
			if updatedUser.Age != 0 {
				users[i].Age = updatedUser.Age
			}
			return users[i], true
		}
	}
	return models.User{}, false
}
