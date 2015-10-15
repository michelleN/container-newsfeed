package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type GithubIssue struct {
	Url       string `json:"html_url"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Number    int    `json:"number"`
	State     string `json:"state"`
	GithubId  int    `json:"id"`
	ClosedAt  string `json:"closed_at"`
	CreatedAt string `json:"create_at"`
	UpdatedAt string `json:"updated_at"`
}

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

func (c *Client) GetIssues(repo string) ([]*GithubIssue, error) {
	req, _ := http.NewRequest("GET", repo, nil)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", c.oauthToken))
	r, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var issues []*GithubIssue
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}
