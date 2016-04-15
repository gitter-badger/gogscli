package cliactions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codegangsta/cli"
	api "github.com/gogits/go-gogs-client"
)

func IssueAdd(auth Auth) func(*cli.Context) {
	return func(c *cli.Context) {
		repository := c.GlobalString("repository")
		if len(repository) == 0 {
			fmt.Printf("Missing \"repository\" flag\n")
			return
		}

		issue := api.Issue{}
		issue.Title = c.String("title")
		issue.Body = c.String("body")

		if len(issue.Title) == 0 {
			fmt.Printf("Missing \"title\" flag\n")
			return
		}
		if len(issue.Body) == 0 {
			fmt.Printf("Missing \"body\" flag\n")
			return
		}

		data, err := json.Marshal(&issue)
		if err != nil {
			fmt.Printf("Error while generating json: %v\n", err)
			return
		}

		url := fmt.Sprintf("%sapi/v1/repos/%s/issues", auth.URL, repository)
		r, _ := http.NewRequest("POST", url, bytes.NewReader(data))
		r.Header.Add("Content-Type", contentTypeJson)
		r.SetBasicAuth(auth.User, auth.Password)
		client := &http.Client{}
		response, err := client.Do(r)
		if err != nil {
			fmt.Printf("Error doing HTTP request: %v\n", err)
			return
		}
		if response.StatusCode != http.StatusCreated {
			fmt.Printf("Unespected response. Status code: %d\n", response.StatusCode)
			return
		}
		fmt.Printf("Issue created\n")
	}
}
