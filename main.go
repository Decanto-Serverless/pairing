package main

import (
	"pairing/env"
	"pairing/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	baseUrl := env.GetInstance().BaseURL

	r.GET(baseUrl+"/pair", handlers.PairFamilies)

	r.Run(env.GetInstance().Port)
}
