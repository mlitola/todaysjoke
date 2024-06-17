package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const jokeUrl = "https://icanhazdadjoke.com/"

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", jokeUrl, nil)
	if err != nil {
		fmt.Printf("error initializing a new http request: %s\n", err)
		os.Exit(1)
	}

	req.Header.Set("User-Agent", "Todays Joke (https://github.com/mlitola/todaysjoke)")
	req.Header.Set("Accept", "text/plain")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error executing a http request: %s\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("i/o error when reading the response body: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s (%s):\n\n", "Todays joke ", time.DateOnly)
	fmt.Printf("%s\n\n", body)
}
