package cliactions

import (
	"bytes"
	"fmt"
	"net/http"
)

const (
	contentTypeJson = "application/json"
)

type Auth struct {
	URL      string
	User     string
	Password string
}

func newRequestContext(method, url string, data []byte, auth Auth) (r *http.Request, client *http.Client) {
	url = fmt.Sprintf("%s%s", auth.URL, url)
	r, _ = http.NewRequest(method, url, bytes.NewReader(data))
	r.Header.Add("Content-Type", contentTypeJson)
	r.SetBasicAuth(auth.User, auth.Password)
	client = &http.Client{}
	return
}
