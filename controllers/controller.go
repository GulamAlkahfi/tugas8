package controller

import (
			"net/http"
	"tugas8/models"
	"tugas8/connect"
	"github.com/labstack/echo/v4"
)

func CreatePekerjaansController(c echo.Context) error {
	var Pekerjaans models.Pekerjaans
	c.Bind(&Pekerjaans)

	result := config.DB.Create(&Pekerjaans)
	if result.Error != nil {
        return c.JSON(500, result.Error.Error())
    }
	return c.JSON(200, Pekerjaans)
}
func GetPekerjaansController(c echo.Context) error {
	var GetPekerjaans []models.Pekerjaans
	result := config.DB.Find(&GetPekerjaans)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}
	return c.JSON(http.StatusOK, GetPekerjaans)
}
func DeletePekerjaans(c echo.Context) error {
    // Mendapatkan ID dari parameter URL
    id := c.Param("id")
    // Lakukan proses penghapusan data di sini (misalnya, dari basis data)
    // ...
		var DeletePekerjaans []models.Pekerjaans
	if err := config.DB.Where("id = ?", id).First(&DeletePekerjaans).Error; err != nil {
		return c.JSON(404, "Data tidak ditemukan")
	}
	config.DB.Delete(&DeletePekerjaans)

	return c.JSON(200, "Data berhasil dihapus")
}

func UpdatePekerjaans(c echo.Context) error {
	  id := c.Param("id")
    // Mendapatkan ID dari parameter URL
		var updateData models.Pekerjaans
			if err := c.Bind(&updateData); err != nil {
				return c.JSON(400, "Data yang diterima tidak valid")
			}
		var existingData models.Pekerjaans
		if err := config.DB.Where("id = ?", id).First(&existingData).Error; err != nil{
		    return c.JSON(404, "Data tidak ditemukan")
			}
			  // Melakukan pembaruan data
			existingData.Pekerjaan = updateData.Pekerjaan
			config.DB.Save(&existingData)
	return c.JSON(200, "Data berhasil diperbarui")
    // Lakukan proses pembaruan data dalam database
    // ...

}
