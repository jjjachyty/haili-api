package services

import "haili/models"

type UserService struct{}

var userModel models.UserModel

func (UserService) GetUsers(params map[string]interface{}) []models.User {
	return userModel.Find(params)
}
