package main

import (
	"github.com/MYanuarBF/go-restaurant-app/internal/database"
	"github.com/MYanuarBF/go-restaurant-app/internal/delivery/rest"
	mRepo "github.com/MYanuarBF/go-restaurant-app/internal/repository/menu"
	rUsecase "github.com/MYanuarBF/go-restaurant-app/internal/usecase/resto"
	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password= dbname=go_resto_app sslmode=disable"
)

func main() {
	e := echo.New()

	db := database.GetDB(dbAddress)
	menuRepo := mRepo.GetRepository(db)

	restoUsecase := rUsecase.GetUsecase(menuRepo)

	h := rest.NewHandler(restoUsecase)
	rest.LoudRoutes(e, h)
	e.Logger.Fatal(e.Start(":14045"))
}
