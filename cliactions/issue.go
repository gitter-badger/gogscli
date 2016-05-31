package cliactions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/andreynering/gogscli/config"
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

func IssueAdd() cli.ActionFunc {
	return func(c *cli.Context) error {
		cfg := config.MustGet()
		if c.NArg() < 2 {
			fmt.Printf("Not enought parameters\n")
			return nil
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
			return nil
		}

		url := fmt.Sprintf("api/v1/repos/%s/issues", repository)
		r, client := newRequestContext("POST", url, data, cfg)
		response, err := client.Do(r)
		if err != nil {
			fmt.Printf("Error doing HTTP request: %v\n", err)
			return nil
		}
		if response.StatusCode != http.StatusCreated {
			content, _ := ioutil.ReadAll(response.Body)
			fmt.Printf("Unespected response. Status code: %d\n%s\n", response.StatusCode, content)
			return nil
		}
		fmt.Printf("Issue created\n")
		return nil
	}
}

func IssueList() cli.ActionFunc {
	return func(c *cli.Context) error {
		cfg := config.MustGet()
		if c.NArg() < 1 {
			fmt.Printf("Not enought parameters\n")
			return nil
		}
		repository := c.Args()[0]

		url := fmt.Sprintf("api/v1/repos/%s/issues", repository)
		r, client := newRequestContext("GET", url, nil, cfg)
		response, err := client.Do(r)
		if err != nil {
			fmt.Printf("Error doing HTTP request: %v\n", err)
			return nil
		}
		if response.StatusCode != http.StatusOK {
			content, _ := ioutil.ReadAll(response.Body)
			fmt.Printf("Unespected response. Status code: %d\n%s\n", response.StatusCode, content)
			return nil
		}

		var issues []api.Issue
		bytes, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(bytes, &issues)
		if err != nil {
			fmt.Printf("Error parsing response: %v\n", err)
			return nil
		}

		if len(issues) == 0 {
			fmt.Printf("No issue\n")
			return nil
		}
		for _, issue := range issues {
			fmt.Printf("#%d %s\n", issue.Index, issue.Title)
		}
		return nil
	}
}
