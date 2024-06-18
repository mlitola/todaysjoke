package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const jokeUrl = "https://icanhazdadjoke.com/"

const userAgent = "Todays Joke (https://github.com/mlitola/todaysjoke)"

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", jokeUrl, nil)
	if err != nil {
		fmt.Printf("Error initializing a new http request: %s\n", err)
		os.Exit(1)
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "text/plain")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making a http request: %s\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("I/O error while reading the response body: %s\n", err)
		os.Exit(1)
	}

	currentTime := time.Now()
	fmt.Printf("%s (%s):\n\n", "Today's joke", currentTime.Format("02.01.2006"))
	fmt.Printf("%s\n\n", body)
}
