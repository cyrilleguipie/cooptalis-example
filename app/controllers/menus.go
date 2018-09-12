package controllers

import (
	"cooptalis-example/app/models"
	"cooptalis-example/app/utils"
	"github.com/revel/revel"
)

type Menus struct {
	*revel.Controller
}

func (c Menus) List(k string, sF string, sD string) revel.Result {
	menus := make([]models.Menu, 0)
	data := models.SearchMenu(menus, k, sF, sD)
	var response utils.JsonResponse
	response.Data = data["objects"]
	response.Success = true
	return c.RenderJSON(response)
}
