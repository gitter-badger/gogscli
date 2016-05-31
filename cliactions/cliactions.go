package cliactions

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/andreynering/gogscli/config"
)

const (
	contentTypeJSON = "application/json"
)

func newRequestContext(method, url string, data []byte, cfg config.Config) (r *http.Request, client *http.Client) {
	url = fmt.Sprintf("%s%s", cfg.Remote.URL, url)
	r, _ = http.NewRequest(method, url, bytes.NewReader(data))
	r.Header.Add("Content-Type", contentTypeJSON)
	r.Header.Add("Authorization", fmt.Sprintf("token %s", cfg.Auth.Token))
	client = &http.Client{}
	return
}
