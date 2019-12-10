package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Summoner represents a response from the Riot API Summoner endpoint
type Summoner struct {
	ProfileIconID int
	Name          string
	Puuid         string
	SummonerLevel int
	AccountID     string
	ID            string
	RevisionDate  uint64
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	riotAPIKey := os.Getenv("RIOT_API_KEY")

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/lol/summoner/v4/summoners/by-name/:name", func(c *gin.Context) {
		summonerName := url.PathEscape(c.Param("name"))
		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s", summonerName), nil)
		if err != nil {
			fmt.Printf("error creating request object: %v", err)
			c.JSON(500, gin.H{
				"error": err,
			})
		}

		req.Header.Add("X-Riot-Token", riotAPIKey)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("error making request: %v", err)
			c.JSON(500, gin.H{
				"error": err,
			})
		}

		defer resp.Body.Close()

		summoner := Summoner{}
		json.NewDecoder(resp.Body).Decode(&summoner)

		c.JSON(200, summoner)
	})

	router.Run()
}
