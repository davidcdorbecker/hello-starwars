package services

import (
	"fmt"
	"hello-starwars/domain"
	"hello-starwars/providers"
	"hello-starwars/utils"
)

type challengeService struct{}

type challlengeServiceInterface interface {
	SayHelloFromSWAPI(id int) (*domain.HelloMessage, utils.Error)
}

var (
	ChallengeService challlengeServiceInterface = &challengeService{}
)

func (cs *challengeService) SayHelloFromSWAPI(id int) (*domain.HelloMessage, utils.Error) {
	person, err := providers.GetPerson(id)
	if err != nil {
		return nil, err
	}

	planet, err := providers.GetPlanet(person.Planet)
	if err != nil {
		return nil, err
	}

	response := &domain.HelloMessage{
		Name:    person.Name,
		Planet:  *planet,
		Message: fmt.Sprintf("Hello!! I am %s and I'm from %s, nice to meet you!", person.Name, planet.Name),
	}

	return response, nil
}
