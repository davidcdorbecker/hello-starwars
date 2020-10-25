package providers

import (
	"github.com/federicoleon/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestGetPersonNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://swapi.dev/api/people/1/",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody: `{
    "name": "Luke Skywalker",
    "height": "172",
    "mass": "77",
    "hair_color": "blond",
    "skin_color": "fair",
    "eye_color": "blue",
    "birth_year": "19BBY",
    "gender": "male",
    "homeworld": "http://swapi.dev/api/planets/1/",
    "films": [
        "http://swapi.dev/api/films/1/",
        "http://swapi.dev/api/films/2/",
        "http://swapi.dev/api/films/3/",
        "http://swapi.dev/api/films/6/"
    ],
    "species": [],
    "vehicles": [
        "http://swapi.dev/api/vehicles/14/",
        "http://swapi.dev/api/vehicles/30/"
    ],
    "starships": [
        "http://swapi.dev/api/starships/12/",
        "http://swapi.dev/api/starships/22/"
    ],
    "created": "2014-12-09T13:50:51.644000Z",
    "edited": "2014-12-20T21:17:56.891000Z",
    "url": "http://swapi.dev/api/people/1/"
}`,
	})

	person, err := GetPerson(1)

	assert.NotNil(t, person)
	assert.Nil(t, err)
	assert.EqualValues(t, "Luke Skywalker", person.Name)
	assert.EqualValues(t, "http://swapi.dev/api/planets/1/", person.Planet)
}

func TestGetPersonNotFound(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://swapi.dev/api/people/1/",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{
    "detail": "Not found"
}`,
	})

	person, err := GetPerson(1)

	assert.Nil(t, person)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.CodeStatus())
	assert.EqualValues(t, "error in the external api", err.Message())
}

func TestGetPersonInvalidModel(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://swapi.dev/api/people/1/",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody: `{
    "name": 999,
	"homeworld": 999,
}`,
	})

	person, err := GetPerson(1)

	assert.Nil(t, person)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.CodeStatus())
	assert.EqualValues(t, "error when trying to unmarshall person", err.Message())
}

func TestGetPlanetNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://swapi.dev/api/planets/1/",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody: `{
    "name": "Tatooine",
    "rotation_period": "23",
    "orbital_period": "304",
    "diameter": "10465",
    "climate": "arid",
    "gravity": "1 standard",
    "terrain": "desert",
    "surface_water": "1",
    "population": "200000",
    "residents": [
        "http://swapi.dev/api/people/1/",
        "http://swapi.dev/api/people/2/",
        "http://swapi.dev/api/people/4/",
        "http://swapi.dev/api/people/6/",
        "http://swapi.dev/api/people/7/",
        "http://swapi.dev/api/people/8/",
        "http://swapi.dev/api/people/9/",
        "http://swapi.dev/api/people/11/",
        "http://swapi.dev/api/people/43/",
        "http://swapi.dev/api/people/62/"
    ],
    "films": [
        "http://swapi.dev/api/films/1/",
        "http://swapi.dev/api/films/3/",
        "http://swapi.dev/api/films/4/",
        "http://swapi.dev/api/films/5/",
        "http://swapi.dev/api/films/6/"
    ],
    "created": "2014-12-09T13:50:49.641000Z",
    "edited": "2014-12-20T20:58:18.411000Z",
    "url": "http://swapi.dev/api/planets/1/"
}`,
	})

	planet, err := GetPlanet("http://swapi.dev/api/planets/1/")

	assert.NotNil(t, planet)
	assert.Nil(t, err)
	assert.EqualValues(t, "Tatooine", planet.Name)
}

func TestGetPlanetNotFound(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://swapi.dev/api/planets/1/",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{
    "detail": "Not found"
}`,
	})

	planet, err := GetPlanet("http://swapi.dev/api/planets/1/")

	assert.Nil(t, planet)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.CodeStatus())
	assert.EqualValues(t, "error in the external api", err.Message())
}

func TestGetPlanetInvalidModel(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://swapi.dev/api/planets/1/",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody: `{
    "name": 999
}`,
	})

	planet, err := GetPlanet("http://swapi.dev/api/planets/1/")

	assert.Nil(t, planet)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.CodeStatus())
	assert.EqualValues(t, "error when trying to unmarshall planet", err.Message())
}
