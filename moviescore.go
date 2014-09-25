package main

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Movie struct {
	name         string
	score        string
	avguserscore string
	genre        string
	cast         string
}

func extractMovie(args []string) (string, error) {
	movie := ""

	if len(args) < 2 {
		return location, errors.New("Please enter a movie search term")
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

func lookupMovies(movie string) []Movie {

}

func main() {
	searchTerm, err := extractMovie(os.Args)

	if err != nil {
		logError(err)
	}

	movies := lookupMovie(searchTerm)

	for _, movie := range movies {
		fmt.Println(movie)
	}
}
