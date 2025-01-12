package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password= dbname=go_resto_app sslmode=disable"
)

type MenuType string

const (
	menuTypeFood  = "food"
	menuTypeDrink = "drink"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
	Type      MenuType
}

func seedDb() {
	foodMenu := []MenuItem{
		{
			Name:      "Bakso Sapi",
			OrderCode: "bs",
			Price:     15000,
			Type:      menuTypeFood,
		}, {
			Name:      "Ayam Goreng",
			OrderCode: "AG",
			Price:     12000,
			Type:      menuTypeFood,
		},
	}
	drinkMenu := []MenuItem{
		{
			Name:      "Wedang Jahe",
			OrderCode: "WJ",
			Price:     5000,
			Type:      menuTypeDrink,
		}, {
			Name:      "Susu Sirup",
			OrderCode: "SS",
			Price:     7000,
			Type:      menuTypeDrink,
		},
	}

	fmt.Println(foodMenu, drinkMenu)

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&MenuItem{})
	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}

func getMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}
	var menuData []MenuItem

	db.Where(MenuItem{Type: MenuType(menuType)}).Find(&menuData)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}

func main() {
	seedDb()
	e := echo.New()
	e.GET("/menu", getMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
