package controllers

import (
	"app/app/models"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

type Products struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c Products) ListProduct() revel.Result {
	var products = []models.Product{
		models.Product{Code: "LK", Price: 1322},
	}
	return c.RenderJSON(products)
}
