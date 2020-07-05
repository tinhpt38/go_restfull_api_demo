package controllers

import (
	"app/app/models"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

type Beers struct {
	*revel.Controller
}

var beers = []models.Beer{
	models.Beer{1, "La Trappe Quadrupel Oak Aged", "Ale", "Bierbrouwerij De Koningshoeven"},
	models.Beer{2, "Cocoa Psycho", "Porter", "BrewDog"},
	models.Beer{3, "American Dream", "Lager", "Mikkeller"},
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c Beers) List() revel.Result {
	return c.RenderJSON(beers)
}
