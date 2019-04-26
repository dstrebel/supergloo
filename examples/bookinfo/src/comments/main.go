package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

func main() {
	port := flag.Uint64("p", 9080, "port")
	failEvenRequests := flag.Bool("fail-half", false, "fail every other request. used to test retries")
	flag.Parse()
	log.Fatal(run(*port, *failEvenRequests))
}

func logReq(r *http.Request) {
	log.Printf("request count: %v", requestCount)
	b, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Printf("error dumping req: %v", err)
	} else {
		log.Printf("request dump:\n%s\n", b)
	}
}

var requestCount int

func run(port uint64, failEvenRequests bool) error {
	m := mux.NewRouter()
	handle(m, "GET", "/health", func(w http.ResponseWriter, r *http.Request) error {
		return json.NewEncoder(w).Encode(map[string]string{"status": "Comments is healthy"})
	})
	handle(m, "GET", "/comments/{id}", func(w http.ResponseWriter, r *http.Request) error {
		requestCount++

		if failEvenRequests && requestCount%2 == 0 {
			return fmt.Errorf("randomly failing request! too bad sucka")
		}

		return replyWithComment(w, r)
	})
	return http.ListenAndServe(fmt.Sprintf(":%v", port), m)
}

func handle(m *mux.Router, method, path string, wrapped func(w http.ResponseWriter, r *http.Request) error) {
	m.Methods(method).Path(path).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logReq(r)

		if err := wrapped(w, r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

type comment struct {
	Author   string `json:"author"`
	Time     string `json:"time"`
	Comment  string `json:"comment"`
	Likes    int    `json:"likes"`
	Dislikes int    `json:"dislikes"`
}

func replyWithComment(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	comment := comments(id)

	return json.NewEncoder(w).Encode(comment)
}

func comments(id string) comment {
	return comment{
		Author:   fmt.Sprintf("lurker %v", id),
		Time:     fmt.Sprintf("%v hours ago", id),
		Likes:    6,
		Dislikes: 3,
		Comment:  randComment(),
	}
}

var commentContents = []string{
	`You're talking to the guy who build his own programming language, his own stdlib, his own game engine, his own C compiler, his own package manager, etc. Does that *really* surprise you? ;)`,
	`“If you try really hard, you *can* write safe code in C/C++” is the flat-earth movement of software engineering`,
	`The use cases for generics are rather limited. Collections, yes. Beyond that, it's usually someone getting cute.`,
	`Learn C programming and the rest will come`,
}

func randComment() string {
	i := rand.Intn(len(commentContents))
	return commentContents[i]
}
