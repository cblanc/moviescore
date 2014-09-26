package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Movie struct {
	Name         string
	Score        string
	Avguserscore string
	Genre        string
	Cast         string
}

type MovieResponse struct {
	Max_pages string
	Count     int
	Results   []Movie
}

func extractMovie(args []string) (string, error) {
	movie := ""

	if len(args) < 2 {
		return movie, errors.New("Please enter a movie search term")
	}

	for i := 1; i < len(args); i++ {
		movie += args[i] + " "
	}

	return strings.Trim(movie, " "), nil
}

func logError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func lookupMovies(movie string) (*MovieResponse, error) {
	var movieResponse MovieResponse

	url := "https://byroredux-metacritic.p.mashape.com/search/movie?max_pages=2&retry=4&title=" + url.QueryEscape(movie)

	req, err := http.NewRequest("POST", url, nil)
	mashape_key := os.Getenv("MASHAPE_KEY")
	req.Header.Add("X-Mashape-Key", mashape_key)

	if err != nil {
		logError(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		logError(err)
	}

	if err := json.NewDecoder(resp.Body).Decode(&movieResponse); err != nil {
		return &movieResponse, err
	}

	return &movieResponse, nil
}

func prettyPrintResults(response *MovieResponse) {
	fmt.Printf("Found %d matching movies\n\n", response.Count)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Movie", "Score", "User Score", "Genre", "Cast"})

	for _, movie := range response.Results {
		data := []string{}
		data = append(data, movie.Name)
		data = append(data, movie.Score)
		data = append(data, movie.Avguserscore)
		data = append(data, movie.Genre)
		data = append(data, movie.Cast)
		table.Append(data)
	}

	table.Render()
}

func main() {
	searchTerm, err := extractMovie(os.Args)

	if err != nil {
		logError(err)
	}

	movies, err := lookupMovies(searchTerm)

	if err != nil {
		logError(err)
	}

	prettyPrintResults(movies)
}
