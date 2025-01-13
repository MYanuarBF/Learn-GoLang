package database

import (
	"fmt"

	"github.com/MYanuarBF/go-restaurant-app/internal/model"
	"github.com/MYanuarBF/go-restaurant-app/internal/model/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	foodMenu := []model.MenuItem{
		{
			Name:      "Bakso Sapi",
			OrderCode: "bs",
			Price:     15000,
			Type:      constant.MenuTypeFood,
		}, {
			Name:      "Ayam Goreng",
			OrderCode: "AG",
			Price:     12000,
			Type:      constant.MenuTypeFood,
		},
	}
	drinkMenu := []model.MenuItem{
		{
			Name:      "Wedang Jahe",
			OrderCode: "WJ",
			Price:     5000,
			Type:      constant.MenuTypeDrink,
		}, {
			Name:      "Susu Sirup",
			OrderCode: "SS",
			Price:     7000,
			Type:      constant.MenuTypeDrink,
		},
	}

	fmt.Println(foodMenu, drinkMenu)

	// db.AutoMigrate(&MenuItem{})
	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}
