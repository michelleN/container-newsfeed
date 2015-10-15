package main

import (
	"encoding/json"
	"fmt"
	//"html"
	"io/ioutil"
	//"log"
	"net/http"
	"os"

	//"github.com/gorilla/mux"
	"gopkg.in/redis.v3"
)

func main() {

	//router := mux.NewRouter().StrictSlash(true)
	redisClient := ExampleNewClient()
	fmt.Println(github("https://api.github.com/repos/deis/deis/issues", redisClient))
	//router.HandleFunc("/", Index)
	//log.Fatal(http.ListenAndServe(":8080", router))
}

//func Index(w http.ResponseWriter, r *http.Request) {
//issues := github("https://api.github.com/repos/deis/deis/issues")
//fmt.Fprintf(w, issues, html.EscapeString(r.URL.Path))
//os.Exit(0)
//}

func check(e error) {
	if e != nil {
		fmt.Print(e.Error())
	}
}

func github(repo string, redisClient *redis.Client) string {
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
	redisClient.Set("time", "githubissue", 0)
	val, _ := redisClient.Get("time").Result()
	fmt.Println("key", val)
	fmt.Println("----")
	fmt.Println(issues[0].Title)
	return issues[0].Url
}

type GithubIssue struct {
	Url   string `json:"html_url"`
	Title string `json:"title"`
	//github_id int    `json:"id"`
}

func ExampleNewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}

//func ExampleClient() {
//err := client.Set("key", "value", 0).Err()
//if err != nil {
//panic(err)
//}

//val, err := client.Get("key").Result()
//if err != nil {
//panic(err)
//}
//fmt.Println("key", val)

//val2, err := client.Get("key2").Result()
//if err == redis.Nil {
//fmt.Println("key2 does not exists")
//} else if err != nil {
//panic(err)
//} else {
//fmt.Println("key2", val2)
//}
//// Output: key value
//// key2 does not exists
//}
