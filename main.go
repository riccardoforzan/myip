package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIResponse struct {
	Address string `json:"ip"`
}

func main() {
	res, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("API error")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var response APIResponse
	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", response.Address)
}
