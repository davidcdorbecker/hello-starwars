package domain

type HelloMessage struct {
	Name    string  `json:"name"`
	Planet  Planets `json:homeworld`
	Message string  `json:"message"`
}
