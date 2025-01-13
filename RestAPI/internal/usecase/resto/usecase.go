package resto

import "github.com/MYanuarBF/go-restaurant-app/internal/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
