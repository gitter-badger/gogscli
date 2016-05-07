package cliactions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codegangsta/cli"
	api "github.com/gogits/go-gogs-client"
)

type issue struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	Assignee  string `json:"assignee"`
	Milestone int    `json:"milestones"`
	Labels    []int  `json:"labels"`
}

func IssueAdd(auth Auth) func(*cli.Context) {
	return func(c *cli.Context) {
		if c.NArg() < 2 {
			fmt.Printf("Not enought parameters\n")
			return
		}
		args := c.Args()
		repository := args[0]

		aissue := issue{}
		aissue.Title = args[1]
		if c.NArg() >= 3 {
			aissue.Body = args[2]
		}

		aissue.Assignee = c.String("assignee")

		data, err := json.Marshal(&aissue)
		if err != nil {
			fmt.Printf("Error while generating json: %v\n", err)
			return
		}

		url := fmt.Sprintf("api/v1/repos/%s/issues", repository)
		r, client := newRequestContext("POST", url, data, auth)
		response, err := client.Do(r)
		if err != nil {
			fmt.Printf("Error doing HTTP request: %v\n", err)
			return
		}
		if response.StatusCode != http.StatusCreated {
			content, _ := ioutil.ReadAll(response.Body)
			fmt.Printf("Unespected response. Status code: %d\n%s\n", response.StatusCode, content)
			return
		}
		fmt.Printf("Issue created\n")
	}
}

func IssueList(auth Auth) func(*cli.Context) {
	return func(c *cli.Context) {
		if c.NArg() < 1 {
			fmt.Printf("Not enought parameters\n")
			return
		}
		repository := c.Args()[0]

		url := fmt.Sprintf("api/v1/repos/%s/issues", repository)
		r, client := newRequestContext("GET", url, nil, auth)
		response, err := client.Do(r)
		if err != nil {
			fmt.Printf("Error doing HTTP request: %v\n", err)
			return
		}
		if response.StatusCode != http.StatusOK {
			content, _ := ioutil.ReadAll(response.Body)
			fmt.Printf("Unespected response. Status code: %d\n%s\n", response.StatusCode, content)
			return
		}

		issues := make([]api.Issue, 0)
		bytes, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(bytes, &issues)
		if err != nil {
			fmt.Printf("Error parsing response: %v\n", err)
			return
		}

		if len(issues) == 0 {
			fmt.Printf("No issue\n")
			return
		}
		for _, issue := range issues {
			fmt.Printf("#%d %s\n", issue.Index, issue.Title)
		}
	}
}
