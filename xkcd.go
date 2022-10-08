package gokcd

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type XKCD struct {
	Number     int    `json:"num"`
	Title      string `json:"title"`
	SafeTitle  string `json:"safe_title"`
	Alt        string `json:"alt"`
	ImageURL   string `json:"img"`
	Link       string `json:"link"`
	News       string `json:"news"`
	Transcript string `json:"transcript"`
	Year       string `json:"year"`
	Month      string `json:"month"`
	Day        string `json:"day"`
}

// FromNum creates a new XKCD from the given comic number/ID.
func FromNum(num int) XKCD {
	c := http.Client{Timeout: time.Second * 2}
	
	// Create the request
	req, err := http.NewRequest("GET", "https://xkcd.com/" + strconv.Itoa(num) + "/info.0.json", nil)
	if err != nil {
		log.Fatal(err)
	}
	
	// Perform the request and save response to res
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	
	// Close the body if it exists
	if res.Body != nil {
		defer res.Body.Close()
	}
	
	// Read the body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	// Unmarshal the JSON response to an XKCD object
	var xkcd XKCD
	err = json.Unmarshal(body, &xkcd)
	if err != nil {
		log.Fatal(err)
	}

	return xkcd
}
