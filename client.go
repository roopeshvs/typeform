package gotypeform

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	apiUrl        = "https://api.typeform.com"
	formsUrl      = "/forms"
	workspacesUrl = "/workspaces"
	accountsUrl   = "/accounts"
	imagesUrl     = "/images"
	userUrl       = "/me"
	themesUrl     = "/themes"
)

type Typeform struct {
	token  string
	client *http.Client
}

func TypeformClient(token string) *Typeform {
	client := &http.Client{}
	return &Typeform{
		token:  token,
		client: client,
	}
}

func (tf *Typeform) buildAndExecRequest(method string, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, apiUrl+url, body)
	req.Header.Add("Accept", "application/json")
	bearer := "Bearer " + tf.token
	req.Header.Add("Authorization", bearer)
	if err != nil {
		panic("Error while building Typeform request")
	}
	resp, err := tf.client.Do(req)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("unknown error, status code: %d", resp.StatusCode)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	return contents, err
}
