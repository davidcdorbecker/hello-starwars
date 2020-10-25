package providers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/federicoleon/golang-restclient/rest"
	"hello-starwars/domain"
	"hello-starwars/helpers"
)

const (
	apiPerson = "https://swapi.dev/api/people/%d/"
)

func GetPerson(id int) (*domain.People, error) {
	response := rest.Get(fmt.Sprintf(apiPerson, id))
	if response == nil || response.Response == nil {
		return nil, errors.New("something fail")
	}

	if response.StatusCode > 299 {
		return nil, errors.New("invalid interface")
	}

	var result domain.People
	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		return nil, errors.New("invalid something.. (TODO)")
	}

	return &result, nil
}

func GetPlanet(planetEndpoint string) (*domain.Planets, error) {

	//fixing a little error from api
	apiEndpointFixed := helpers.RedirectHttps(planetEndpoint)
	response := rest.Get(apiEndpointFixed)
	if response == nil || response.Response == nil {
		return nil, errors.New("something fail")
	}

	if response.StatusCode > 299 {
		return nil, errors.New("invalid interface")
	}

	var result domain.Planets
	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		return nil, errors.New("invalid something.. (TODO)")
	}

	return &result, nil
}
