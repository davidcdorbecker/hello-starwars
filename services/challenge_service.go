package services

import (
	"fmt"
	"hello-starwars/domain"
	"hello-starwars/providers"
	"hello-starwars/utils"
)

var (
	ChallengeService ChalllengeServiceInterface = &challengeService{}
)

type challengeService struct{}

type ChalllengeServiceInterface interface {
	SayHelloFromSWAPI(id int) (*domain.HelloMessage, utils.Error)
}

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
