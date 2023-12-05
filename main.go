package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Seat struct {
	ID            int     `json:"id"`
	Number        string  `json:"number"`
	Occupied      bool    `json:"occupied"`
	PassengerName string  `json:"passengerName,omitempty"`
	PassengerDOB  string  `json:"passengerDOB,omitempty"`
	PassengerID   string  `json:"passengerID,omitempty"`
	Price         float64 `json:"price,omitempty"`
}

type SeatResponse struct {
	ID            int     `json:"id"`
	Number        string  `json:"number"`
	Occupied      bool    `json:"occupied"`
	PassengerName string  `json:"passengerName,omitempty"`
	PassengerDOB  string  `json:"passengerDOB,omitempty"`
	PassengerCPF  string  `json:"passengerCPF,omitempty"`
	Price         float64 `json:"price,omitempty"`
}

type Flight struct {
	ID          int     `json:"id"`
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	Departure   string  `json:"departure"`
	Seats       []Seat  `json:"seats"`
	Price       float64 `json:"price,omitempty"` // Preço único para os assentos deste voo
}

var flights = []Flight{
	{ID: 1, Origin: "São Paulo", Destination: "Rio de Janeiro", Departure: "20/12/2023 10:00", Seats: generateSeats(20, 500.00), Price: 500.00},
	{ID: 2, Origin: "Brasília", Destination: "Salvador", Departure: "20/12/2023 10:30", Seats: generateSeats(20, 650.00), Price: 650.00},
	{ID: 3, Origin: "Belo Horizonte", Destination: "Fortaleza", Departure: "20/12/2023 11:00", Seats: generateSeats(20, 400.00), Price: 400.00},
	{ID: 4, Origin: "Recife", Destination: "Porto Alegre", Departure: "20/12/2023 11:30", Seats: generateSeats(20, 1200.00), Price: 1200.00},
	{ID: 5, Origin: "Manaus", Destination: "Curitiba", Departure: "20/12/2023 12:00", Seats: generateSeats(20, 750.00), Price: 750.00},
	{ID: 6, Origin: "Belém", Destination: "Goiânia", Departure: "20/12/2023 12:30", Seats: generateSeats(20, 480.00), Price: 480.00},
	{ID: 7, Origin: "Campinas", Destination: "São Luís", Departure: "20/12/2023 13:00", Seats: generateSeats(20, 820.00), Price: 820.00},
	{ID: 8, Origin: "Maceió", Destination: "Florianópolis", Departure: "20/12/2023 13:30", Seats: generateSeats(20, 1000.00), Price: 1000.00},
	{ID: 9, Origin: "Natal", Destination: "Cuiabá", Departure: "20/12/2023 14:00", Seats: generateSeats(20, 740.00), Price: 740.00},
	{ID: 10, Origin: "João Pessoa", Destination: "Vitória", Departure: "20/12/2023 14:30", Seats: generateSeats(20, 560.00), Price: 560.00},
}

func generateSeats(numSeats int, price float64) []Seat {
	seats := make([]Seat, numSeats)
	for i := 1; i <= numSeats; i++ {
		seats[i-1] = Seat{ID: i, Number: fmt.Sprintf("A%d", i), Occupied: false, Price: price}
	}
	return seats
}

func updateSeatHandler(c *gin.Context) {
	flightID, err := strconv.Atoi(c.Param("flightID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de voo inválido"})
		return
	}

	seatNumber := c.Param("seatNumber")

	flight, err := findFlightByID(flightID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Voo não encontrado"})
		return
	}

	// Encontra o assento com base no número do assento
	var updatedSeat Seat
	for i, seat := range flight.Seats {
		if seat.Number == seatNumber {
			// Atualiza o assento com os dados da requisição
			if err := c.ShouldBindJSON(&updatedSeat); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON inválido"})
				return
			}

			flight.Seats[i] = updatedSeat

			c.JSON(http.StatusOK, updatedSeat)
			return
		}
	}

	// Se o assento não for encontrado, retorna 404
	c.JSON(http.StatusNotFound, gin.H{"error": "Assento não encontrado"})
}

// Adicione esta estrutura para os detalhes do voo
type FlightDetails struct {
	ID          int     `json:"id"`
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	Departure   string  `json:"departure"`
	Seats       []Seat  `json:"seats"`
	Price       float64 `json:"price,omitempty"`
}

// Modifique a função getFlightDetailsHandler
func getFlightDetailsHandler(c *gin.Context) {
	flightID, err := strconv.Atoi(c.Param("flightID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de voo inválido"})
		return
	}

	flight, err := findFlightByID(flightID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Voo não encontrado"})
		return
	}

	flightDetails := FlightDetails{
		ID:          flight.ID,
		Origin:      flight.Origin,
		Destination: flight.Destination,
		Departure:   flight.Departure,
		Seats:       flight.Seats,
		Price:       flight.Price,
	}

	c.JSON(http.StatusOK, flightDetails)
}

func findFlightByID(id int) (*Flight, error) {
	for _, flight := range flights {
		if flight.ID == id {
			return &flight, nil
		}
	}
	return nil, fmt.Errorf("Voo não encontrado")
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	// Adicionar 20 assentos para cada um dos voos restantes (1 a 10)
	for i := 1; i <= 10; i++ {
		flights[i-1].Seats = generateSeats(20, flights[i-1].Price)
	}

	// Endpoint para obter voos
	router.GET("/flights", func(c *gin.Context) {
		c.JSON(http.StatusOK, flights)
	})

	// Endpoint para obter assentos disponíveis no voo
	router.GET("/flights/:flightID/seats", func(c *gin.Context) {
		flightID, err := strconv.Atoi(c.Param("flightID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de voo inválido"})
			return
		}

		flight, err := findFlightByID(flightID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Voo não encontrado"})
			return
		}

		c.JSON(http.StatusOK, flight.Seats)
	})

	// Endpoint para atualizar o assento
	router.PATCH("/flights/:flightID/seats/:seatNumber", updateSeatHandler)
	router.GET("/flights/:flightID/details", getFlightDetailsHandler)

	router.Run(":8080")
}
