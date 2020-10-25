package providers

import (
	"encoding/json"
	"fmt"
	"github.com/federicoleon/golang-restclient/rest"
	"hello-starwars/domain"
	"hello-starwars/helpers"
	"hello-starwars/utils"
	"net/http"
)

const (
	apiPerson = "https://swapi.dev/api/people/%d/"
)

func GetPerson(id int) (*domain.People, utils.Error) {
	response := rest.Get(fmt.Sprintf(apiPerson, id))
	if response == nil || response.Response == nil {
		return nil, utils.NewRestError("invalid id param", http.StatusInternalServerError)
	}

	if response.StatusCode > 299 {
		return nil, utils.NewRestError("error in the external api", http.StatusInternalServerError)
	}

	var result domain.People
	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		return nil, utils.NewRestError("error when trying to unmarshall person", http.StatusInternalServerError)
	}

	return &result, nil
}

func GetPlanet(planetEndpoint string) (*domain.Planets, utils.Error) {

	//fixing a little error from api
	apiEndpointFixed := helpers.RedirectHttps(planetEndpoint)
	response := rest.Get(apiEndpointFixed)
	if response == nil || response.Response == nil {
		return nil, utils.NewRestError("invalid id param", http.StatusInternalServerError)
	}

	if response.StatusCode > 299 {
		return nil, utils.NewRestError("error in the external api", http.StatusInternalServerError)
	}

	var result domain.Planets
	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		return nil, utils.NewRestError("error when trying to unmarshall planet", http.StatusInternalServerError)
	}

	return &result, nil
}
