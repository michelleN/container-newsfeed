package github

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Client struct {
	client     *http.Client
	oauthToken string
}

func NewClient() *Client {
	return &Client{
		client:     &http.Client{},
		oauthToken: os.Getenv("OAUTH_TOKEN"),
	}
}

func (c *Client) Get(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", c.oauthToken))
	r, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
