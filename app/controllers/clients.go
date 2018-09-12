package controllers

import (
	"cooptalis-example/app/models"
	"cooptalis-example/app/utils"
	"github.com/revel/revel"
)

type Clients struct {
	*revel.Controller
}

func (c Clients) List(k string, sF string, sD string) revel.Result {
	clients := make([]models.Client, 0)
	data := models.SearchClient(clients, k, sF, sD)
	var response utils.JsonResponse
	response.Data = data["objects"]
	response.Success = true

	return c.RenderJSON(response)
}
