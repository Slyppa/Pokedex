package Pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type ResourceList struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func Locations(page int) {
	offset := int((page - 1) * 20)

	basePath := "https://pokeapi.co/api/v2/location-area"
	u, err := url.Parse(basePath)
	if err != nil {
		fmt.Println("Error parsing base URL:", err)
		return
	}

	q := u.Query()
	q.Set("limit", "20")
	q.Set("offset", strconv.Itoa(offset))
	u.RawQuery = q.Encode()
	fmt.Println("Requesting URL:", u.String())

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("Network Error:", err)
		return
	}
	defer resp.Body.Close()

	var data ResourceList
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("JSON decode error:", err)
		return
	}

	fmt.Println("\n--- Location Areas (Page 1) ---")
	for i, area := range data.Results {
		fmt.Printf("%2d. %s\n", i+1+offset, area.Name)
	}
}
