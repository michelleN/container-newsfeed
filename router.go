package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	issues := github("https://api.github.com/repos/deis/deis/issues")
	fmt.Fprintf(w, issues, html.EscapeString(r.URL.Path))
	os.Exit(0)
}

func check(e error) {
	if e != nil {
		fmt.Print(e.Error())
	}
}

func github(repo string) string {
	client := &http.Client{}
	oauthToken := os.Getenv("OAUTH_TOKEN")
	req, _ := http.NewRequest("GET", repo, nil)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", oauthToken))
	r, e := client.Do(req)
	check(e)
	defer r.Body.Close()

	var issues []*GithubIssue
	body, e := ioutil.ReadAll(r.Body)
	check(e)
	err := json.Unmarshal(body, &issues)
	check(err)
	fmt.Println(issues[0].Title)
	return issues[0].Url
}

type GithubIssue struct {
	Url   string `json:"html_url"`
	Title string `json:"title"`
	//github_id int    `json:"id"`
}
