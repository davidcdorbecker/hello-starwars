package services

import (
	"fmt"
	"hello-starwars/domain"
	"hello-starwars/providers"
)

var (
	ChallengeService ChalllengeServiceInterface = &challengeService{}
)

type challengeService struct{}

type ChalllengeServiceInterface interface {
	HelloStarwars(id int) (*domain.HelloMessage, error)
}

func (cs *challengeService) HelloStarwars(id int) (*domain.HelloMessage, error) {
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
