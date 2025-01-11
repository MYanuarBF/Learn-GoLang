package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type menuItems struct {
	Name      string
	OrderCode string
	Price     int
}

func getFoodMenu(c echo.Context) error {
	foodMenu := []menuItems{
		{
			Name:      "Bakso Sapi",
			OrderCode: "bs",
			Price:     15000,
		}, {
			Name:      "Ayam Goreng",
			OrderCode: "AG",
			Price:     12000,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}

func getDrinkMenu(c echo.Context) error {
	drinkMenu := []menuItems{
		{
			Name:      "Wedang Jahe",
			OrderCode: "WJ",
			Price:     5000,
		}, {
			Name:      "Susu Sirup",
			OrderCode: "SS",
			Price:     7000,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": drinkMenu,
	})
}

func main() {
	e := echo.New()
	e.GET("menu/food", getFoodMenu)
	e.GET("menu/drink", getDrinkMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
