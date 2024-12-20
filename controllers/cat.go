package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
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

type FavoriteRequest struct {
	ImageID string `json:"image_id"`
	SubID   string `json:"sub_id"`
}

func (c *CatController) AddToFavorites() {
	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		c.Ctx.WriteString("Error loading .env file")
		return
	}

	apiKey := os.Getenv("API_Key")
	if apiKey == "" {
		c.Ctx.WriteString("API Key is missing")
		return
	}

	// Get the request body
	var favoriteRequest FavoriteRequest
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &favoriteRequest)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Ctx.WriteString(fmt.Sprintf("Error parsing JSON: %s", err.Error()))
		return
	}

	// Forward the request to The Cat API
	client := &http.Client{}
	url := "https://api.thecatapi.com/v1/favourites"
	reqBody, err := json.Marshal(map[string]string{
		"image_id": favoriteRequest.ImageID,
		"sub_id":   favoriteRequest.SubID,
	})

	if err != nil {
		c.Ctx.WriteString("Error creating request body")
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		c.Ctx.WriteString("Error creating request")
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

	// Send the request to The Cat API
	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.WriteString("Error sending request to The Cat API")
		return
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode == http.StatusOK {
		c.Ctx.WriteString("Cat added to favorites successfully!")
	} else {
		c.Ctx.WriteString("Failed to add cat to favorites")
	}
}

func (c *CatController) GetFavorites() {
	// Load environment variables (API key)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_Key")
	if apiKey == "" {
		c.Ctx.WriteString("API Key is missing")
		return
	}

	// Prepare the API request to get the user's favorites
	req, err := http.NewRequest("GET", "https://api.thecatapi.com/v1/favourites", nil)
	if err != nil {
		c.Ctx.WriteString("Error preparing the request")
		return
	}

	// Set the API Key header
	req.Header.Set("x-api-key", apiKey)

	// Make the API call
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.WriteString("Error fetching favorites")
		return
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Ctx.WriteString("Error reading response")
		return
	}

	var favorites []struct {
		Image struct {
			ID  string `json:"id"`
			URL string `json:"url"`
		} `json:"image"`
	}
	err = json.Unmarshal(body, &favorites)
	if err != nil {
		c.Ctx.WriteString("Error unmarshalling JSON")
		return
	}

	// Pass the favorite images to the template
	if len(favorites) > 0 {
		c.Data["Favorites"] = favorites
	} else {
		c.Data["Favorites"] = nil
	}

	// Render the template with both the random cat and favorite sections
	c.TplName = "cat.tpl"
}
