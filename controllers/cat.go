package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type CatController struct {
	web.Controller
}

type CatResponse struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// GetCatImage is responsible for getting a random cat image
func (c *CatController) GetCatImage() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_Key")
	fmt.Println("API Key: ", os.Getenv("API_Key"))
	if apiKey == "" {
		c.Ctx.WriteString("API Key is missing")
		return
	}
	c.Data["API_Key"] = apiKey

	fmt.Println("API in c DaTa", c.Data["API_Key"])

	// Make a GET request to The Cat API
	response, err := http.Get("https://api.thecatapi.com/v1/images/search?size=med&mime_types=jpg&format=json&has_breeds=true&order=RANDOM&page=0&limit=1")
	if err != nil {
		c.Ctx.WriteString("Error fetching cat image")
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.Ctx.WriteString("Error reading response body")
		return
	}

	var data []CatResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		c.Ctx.WriteString("Error unmarshalling JSON")
		return
	}

	// Check if we have a valid response
	if len(data) > 0 {
		c.Data["CatImageURL"] = data[0].URL
		c.Data["CatImageID"] = data[0].ID
	} else {
		c.Data["CatImageURL"] = "No cat image found"
		c.Data["CatImageID"] = ""
	}
	c.TplName = "cat.tpl"
	c.Render()
}

type VoteRequest struct {
	ImageID string `json:"image_id"`
	SubID   string `json:"sub_id"`
	Value   int    `json:"value"`
}

func (c *CatController) Vote() {
	// Read the JSON body
	var voteRequest VoteRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &voteRequest)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Ctx.WriteString(fmt.Sprintf("Error parsing JSON: %s", err.Error()))
		return
	}

	// Log the received data for debugging
	fmt.Println("Received Vote:", voteRequest)

	// Check if the value is valid (1 for upvote, 2 for downvote)
	if voteRequest.Value != 1 && voteRequest.Value != 2 {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Ctx.WriteString("Invalid vote value")
		return
	}

	// Respond with a success message
	c.Ctx.WriteString("Vote recorded successfully")
}
