package controller

import (
	"tugas8/middleware"
	"tugas8/models"
	"tugas8/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

/*func LoginController(c echo.Context) error {
	return c.String(200, "Hello World")
}*/

func RegisterController(c echo.Context) error {
	var userRegister models.User
	c.Bind(&userRegister)

	result, _ := bcrypt.GenerateFromPassword(
		[]byte(userRegister.Password),
		bcrypt.DefaultCost,
	)
	userRegister.Password = string(result)
	err := repository.Register(&userRegister)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Baseresponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	var userResponse models.UserResponse
	userResponse.Id = userRegister.Id
	userResponse.Email = userRegister.Email
	userResponse.Token = middlewares.GenerateJWTToken(
		int(userResponse.Id), //dipaksa ke int karena uInt di structnya
		userResponse.Name,
	)

	return c.JSON(http.StatusOK, models.Baseresponse{
		Message: "Berhasil Registered Successfully",
		Status:  true,
		Data:    userResponse,
	})
}
