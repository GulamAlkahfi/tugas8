package controller

import (
			"net/http"
	"tugas8/models"
	"tugas8/connect"
	"github.com/labstack/echo/v4"
)
func CreateWargasController(c echo.Context) error {
	var Wargas models.Wargas
	c.Bind(&Wargas)

	result := config.DB.Create(&Wargas)
	if result.Error != nil {
        return c.JSON(500, result.Error.Error())
    }
	return c.JSON(200, Wargas)
}

func GetWargasController(c echo.Context) error {
	var Wargas []models.Wargas
	result := config.DB.Find(&Wargas)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}
	return c.JSON(http.StatusOK, Wargas)
}
