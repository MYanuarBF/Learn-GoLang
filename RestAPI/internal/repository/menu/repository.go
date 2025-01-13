package menu

import "github.com/MYanuarBF/go-restaurant-app/internal/model"

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
