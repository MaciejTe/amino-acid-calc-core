package usda

//package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

const usdaURL string = "api.nal.usda.gov/fdc/"

// USDAClient structure allows to communicate with USDA REST API.
type USDAClient struct {
	apiKey string
	client *resty.Client
}

// NewClient creates new USDAClient object.
func NewClient(apiKey string) *USDAClient {
	restyClient := resty.New()
	return &USDAClient{apiKey: apiKey, client: restyClient}
}

// IngredientSearch endpoint returns a list of foods that match the search criteria.
func (u USDAClient) FoodSearch(food string) (*resty.Response, error) {
	body := fmt.Sprintf(`{
		"generalSearchInput":"%s", 
		"pageNumber":"1", 
		"includeDataTypes": {
			"Branded": true,
			"Survey (FNDDS)": true,
			"Foundation": true
		}
	}`, food)
	endpoint := fmt.Sprintf("https://%s@%s/v1/search", u.apiKey, usdaURL)
	return u.client.R().SetHeader("Content-Type", "application/json").SetBody(body).Post(endpoint)
}

// IngredientDetails endpoint provides all available details on a particular food.
// id - unique identifier for the food.
func (u USDAClient) FoodDetails(id string) (*resty.Response, error) {
	endpoint := fmt.Sprintf("https://%s@%s/v1/%s", u.apiKey, usdaURL, id)
	return u.client.R().SetHeader("Content-Type", "application/json").Get(endpoint)
}
