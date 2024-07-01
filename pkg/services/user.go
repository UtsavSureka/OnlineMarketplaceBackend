package services

import (
	userdb "Ecomm/pkg/db/UserDB"
	"Ecomm/pkg/models"
	"Ecomm/pkg/utils"
	"errors"
	"fmt"
)

func RegisterNewUser(newUser models.User) (int, error) {

	users, err := userdb.GetAllUsers()
	if err != nil {
		return 0, err
	}

	for _, user := range users {
		if user.User_name == newUser.User_name || user.Email == newUser.Email {
			return 0, errors.New("user name or email already exists")
		}
	}

	id, err := userdb.AddNewUser(newUser)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func UserLogin(userName, password string) (string, error) {

	Users, err := userdb.GetAllUsers()
	if err != nil {
		return "", err
	}

	var userFound bool
	var passwordHash string

	var authenticatedUser models.User

	// Check if user found
	for _, user := range Users {
		if user.User_name == userName {
			userFound = true
			authenticatedUser = user
			passwordHash = user.Password
			break
		}
	}
	if !userFound {
		return "", errors.New("user not found")
	}

	//Check if password hash and password matches
	CheckPasswordHash, err := utils.CheckPasswordHash(password, passwordHash)
	if !CheckPasswordHash && err != nil {
		return "", errors.New("incorrect UserName or Password")
	}

	//Generate jwt token

	token, err := utils.GenerateJwtToken(authenticatedUser)
	if err != nil {
		return "", errors.New("error generating JWT Token")
	}

	return token, nil
}

func GetUserById(userId int) (models.User, error) {
	Users, err := userdb.GetAllUsers()

	if err != nil {
		return models.User{}, err
	}

	var userFound bool

	var userDetail models.User

	for _, user := range Users {
		if user.Id == userId {
			userDetail = user
			userFound = true
		}
	}

	if !userFound {
		return models.User{}, fmt.Errorf("user with id %d is not present", userId)
	}

	return userDetail, nil

}

func UpdateUserDetails(updatedUser models.User) error {

	Users, err := userdb.GetAllUsers()
	if err != nil {
		return err
	}

	var userFound bool

	for _, user := range Users {
		if user.Id == updatedUser.Id {
			userFound = true
		}
	}

	if !userFound {
		return fmt.Errorf("user with id %d is not present", updatedUser.Id)
	}

	err = userdb.UpdateUserDetails(updatedUser)
	if err != nil {
		return err
	}

	return nil
}
