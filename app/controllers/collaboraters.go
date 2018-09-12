package controllers

import (
	"cooptalis-example/app/models"
	"cooptalis-example/app/utils"
	"fmt"
	"github.com/revel/revel"
)

type Collaboraters struct {
	*revel.Controller
}

func (c Collaboraters) List(k string, sF string, sD string) revel.Result {
	collaboraters := make([]models.Collaborater, 0)
	data := models.SearchCollaborater(collaboraters, k, sF, sD)
	var response utils.JsonResponse
	response.Data = data["objects"]
	response.Success = true

	return c.RenderJSON(response)
}

func (c Collaboraters) Update(id uint) revel.Result {
	var response utils.JsonResponse
	var jsonData map[string]string
	c.Params.BindJSON(&jsonData)
	data, err := models.FindCollaboraterById(id)
	if err != nil {
		response.Success = false
		response.Message = "Database error"

	} else {
		if value, ok := jsonData["country"]; ok {
			data.Country = value
		} else {
			fmt.Println("key country not found")
		}

		if value, ok := jsonData["client"]; ok {
			data.Client = value
		} else {
			fmt.Println("key client not found")
		}
		if value, ok := jsonData["state"]; ok {
			data.State = value
		} else {
			fmt.Println("key state not found")
		}
		models.Gorm.Save(&data)
		response.Success = true
		response.Data = data
	}
	return c.RenderJSON(response)
	//Fields to update : Country, Client, Status

}

func (c Collaboraters) FindById(id uint) revel.Result {
	var response utils.JsonResponse
	var data models.Collaborater
	data, err := models.FindCollaboraterById(id)
	if err != nil {
		response.Success = false
		response.Message = "Database error"
	} else {
		response.Success = true
		response.Data = data
	}
	return c.RenderJSON(response)

}
