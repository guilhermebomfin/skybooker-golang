package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Seat struct {
	ID       int    `json:"id"`
	Number   string `json:"number"`
	Occupied bool   `json:"occupied"`
	Owner    string `json:"owner,omitempty"`
	// Outros campos relevantes
}

type Flight struct {
	ID          int    `json:"id"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Departure   string `json:"departure"`
	// Outros campos relevantes
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	// Endpoint para obter voos
	router.GET("/flights", func(c *gin.Context) {
		flights := []Flight{
			{ID: 1, Origin: "Aeroporto A", Destination: "Aeroporto B", Departure: "20/12/2023 10:00"},
			{ID: 2, Origin: "Aeroporto C", Destination: "Aeroporto D", Departure: "20/12/2023 10:30"},
			// Adicione mais voos conforme necessário
		}

		c.JSON(http.StatusOK, flights)
	})

	// Endpoint para obter assentos disponíveis no voo de ID 1
	router.GET("/flights/1/seats", func(c *gin.Context) {
		seats := []Seat{
			{ID: 1, Number: "A1", Occupied: false},
			{ID: 2, Number: "A2", Occupied: false},
			{ID: 3, Number: "B1", Occupied: true, Owner: "João"},
			// Adicione mais conforme necessário
		}

		c.JSON(http.StatusOK, seats)
	})

	// Endpoint para obter assentos disponíveis no voo de ID 2
	router.GET("/flights/2/seats", func(c *gin.Context) {
		seats := []Seat{
			{ID: 1, Number: "C1", Occupied: false},
			{ID: 2, Number: "C2", Occupied: true, Owner: "Maria"},
			{ID: 3, Number: "D1", Occupied: false},
			// Adicione mais conforme necessário
		}

		c.JSON(http.StatusOK, seats)
	})

	router.Run(":8080")
}
