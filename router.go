package main

import (
	"fmt"
	"html"
	"io/ioutil"
	//"log"
	"encoding/json"
	"net/http"
	"os"

	//"github.com/gorilla/mux"
)

func main() {

	//router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", Index)
	fmt.Println(github())
	//log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	os.Exit(0)
}

func check(e error) {
	if e != nil {
		fmt.Print(e.Error())
	}
}

func github() string {
	client := &http.Client{}
	oauthToken := os.Getenv("OAUTH_TOKEN")
	req, _ := http.NewRequest("GET", "https://api.github.com/repos/deis/deis/issues", nil)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", oauthToken))
	r, e := client.Do(req)
	check(e)
	fmt.Println("1")
	defer r.Body.Close()

	var data GithubResponse
	body, e := ioutil.ReadAll(r.Body)
	fmt.Println("2")
	check(e)
	err := json.Unmarshal(body, &data)
	fmt.Println("3")
	check(err)
	return data.issues[0].url
}

type GithubResponse struct {
	issues map[string]GithubIssue
}

type GithubIssue struct {
	url string `json:"html_url"`
	//github_id int    `json:"id"`
}
