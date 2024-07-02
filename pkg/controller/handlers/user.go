package handlers

import (
	"Ecomm/pkg/models"
	"Ecomm/pkg/services"
	"Ecomm/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserRegistrationHandler registers a new user to the database
// @Summary Register a new user
// @Description Register a new user and add the details in the database
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user body RegisterUserRequest true "User registration request"
// @Success 200 {object} RegisterUserResponse
// @Failure 400 {object} RegisterUserResponse
// @Failure 500 {object} RegisterUserResponse
// @Router users/register [post]
func UserRegistrationHandler(ctx *gin.Context) {

	var req RegisterUserRequest
	var resp RegisterUserResponse

	err := ctx.BindJSON(&req)
	if err != nil {
		resp.Error = err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	hashedPassword, err := utils.HashedPassword(req.Password)

	if err != nil {
		resp.Error = err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	newUser := models.User{
		Id:         req.Id,
		First_name: req.First_name,
		Last_name:  req.Last_name,
		User_name:  req.User_name,
		Password:   hashedPassword,
		Email:      req.Email,
		IsAdmin:    req.IsAdmin,
	}

	lastInsertedId, err := services.RegisterNewUser(newUser)

	if err != nil {
		resp.Error = err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Id = lastInsertedId
	resp.Message = "User added successfully"

	ctx.JSON(http.StatusOK, resp)

}

func UserLoginHandler(ctx *gin.Context) {
	var req UserLoginRequest
	var resp UserLoginResponse

	err := ctx.BindJSON(&req)
	if err != nil {
		resp.Error = err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	token, err := services.UserLogin(req.User_Name, req.Password)

	if err != nil {
		resp.Error = err.Error()
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

	resp.Message = "Login Successfull"
	resp.Token = token
	ctx.JSON(http.StatusOK, resp)

}

func GetUserHandler(ctx *gin.Context) {
	var resp GetUserByIdResponse
	idString := ctx.Query("id")

	if idString == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid User Id",
		})
		return
	}

	user, err := services.GetUserById(id)
	if err != nil {
		resp.Error = err.Error()
		ctx.JSON(http.StatusBadRequest, resp.Error)
		return
	}

	resp = GetUserByIdResponse{
		Id:         user.Id,
		First_name: user.First_name,
		Last_name:  user.Last_name,
		User_name:  user.User_name,
		Email:      user.Email,
	}

	ctx.JSON(http.StatusOK, resp)
}

func UpdateUserDetails(ctx *gin.Context) {
	var req UpdateUserRequest
	var resp UpdateUserResponse

	err := ctx.BindJSON(&req)
	if err != nil {
		resp.Error = err.Error()
		ctx.JSON(http.StatusBadRequest, resp.Error)
		return
	}
	hashedPassword, err := utils.HashedPassword(req.Password)
	if err != nil {
		resp.Error = err.Error()
		ctx.JSON(http.StatusInternalServerError, resp.Error)
		return
	}
	user := models.User{
		Id:         req.Id,
		First_name: req.First_name,
		Last_name:  req.Last_name,
		User_name:  req.User_name,
		Password:   hashedPassword,
		Email:      req.Email,
	}
	err = services.UpdateUserDetails(user)
	if err != nil {
		resp.Error = err.Error()
		ctx.JSON(http.StatusBadRequest, resp.Error)
		return
	}

	resp.Message = "Successfully Updated the user details"
	ctx.JSON(http.StatusOK, resp)
}
