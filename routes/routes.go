package routes

import (
	"echo-api/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Connected struct {
	Status  bool   `json:"status" xml:"status"`
	Code    string `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
}

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/zendy", func(c echo.Context) error {
		C := &Connected{
			Status:  true,
			Code:    "200",
			Message: "Api Connected successfully",
		}
		return c.JSON(http.StatusOK, C)
	})

	e.GET("/pegawai", controllers.FetchAllPegawai)
	e.POST("/pegawai", controllers.StoredPegawai)
	e.PUT("/pegawai", controllers.UpdatePegawai)
	e.DELETE("/pegawai", controllers.DeletePegawai)

	return e
}
