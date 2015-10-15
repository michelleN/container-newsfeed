package github

import (
	"encoding/json"
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

func (c *Client) GetIssues(repo string) ([]*GithubIssue, error) {
	body, err := c.Get(repo)
	if err != nil {
		return nil, err
	}

	var issues []*GithubIssue
	err = json.Unmarshal(body, &issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}
