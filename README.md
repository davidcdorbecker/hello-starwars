# hello-starwars
Brief challenge 

This challenge was develop in Golang v 1.15 using go modules

This API is a brief challenge. It has the following two endpoints: 

## /challenge
This enpoint only display a hello world message.

## /challenge/starwars?id=1
This endpoint will display a custom hellow message from SWAPI. It will display a greeting given an id (valid integer of the SWAPI), for example: 

```javascript
 {
    "name": "Luke Skywalker",
    "Planet": {
        "name": "Tatooine"
    },
    "message": "Hello!! I am Luke Skywalker and I'm from Tatooine, nice to meet you!"
 }

```

To run this API you only have to run this command in src (default port: 8080): 

go run main.go

## Dependencies

#### Libraries used in this project: 
##### github.com/gin-gonic/gin
##### github.com/federicoleon/golang-restclient
##### github.com/stretchr/testify

