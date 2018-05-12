package main

import (
	"time"
	"net/url"
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
	"os"
	"log"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string    `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))

	response, err := http.Get(IssuesURL + "?q=" + q)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", response.Status)
	}

	var result IssuesSearchResult
	if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func formatDate(item time.Time) string {
	return item.Format("02/01/2006")
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		var day string

		if int(time.Since(item.CreatedAt).Hours()/24) == 0 {
			day = fmt.Sprintf("today %v", formatDate(item.CreatedAt))
		} else {
			day = fmt.Sprintf("%d days ago %v", int(time.Since(item.CreatedAt).Hours()/24),
				formatDate(item.CreatedAt))
		}

		fmt.Printf("#%-5d %9.9s %22.21s %.55s \n", item.Number, item.User.Login, day,
			item.Title)

	}
}
