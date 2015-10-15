package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/redis.v3"

	"github.com/adamreese/container-newsfeed/github"
)

var RedisClient *redis.Client
var gh = github.NewClient()

func main() {

	router := mux.NewRouter().StrictSlash(true)
	RedisClient = ExampleNewClient()
	go runGithub()
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	var cursor int64
	for {
		var keys []string
		var err error
		cursor, keys, err = RedisClient.Scan(cursor, "", 10).Result()
		if err != nil {
			panic(err)
		}
		for _, key := range keys {
			val, _ := RedisClient.Get(key).Result()
			fmt.Fprintf(w, "%q\n", val)
		}
		if cursor == 0 {
			break
		}
	}
}

func runGithub() {
	issuesUrl := []string{"https://api.github.com/repos/deis/deis/issues", "https://api.github.com/repos/kubernetes/kubernetes/issues"}
	for _, url := range issuesUrl {
		fmt.Printf("%v\n", url)
		issues, err := gh.GetIssues(url)
		check(err)
		fmt.Printf("%v\n", issues[0])
		RedisClient.Set("time", "githubissue", 0)
	}
}

func check(e error) {
	if e != nil {
		fmt.Print(e.Error())
	}
}

func ExampleNewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}

func githubQueryUrl(url string) {
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
