package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	type user struct {
		Name string
	}

	users := []user{}

	http.HandleFunc("/name", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello i am vikas"))
	})
	http.HandleFunc("/address", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("village chiri rohtak"))
	})
	http.HandleFunc("/profession", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("marketing"))
	})
	http.HandleFunc("/createuser", func(rw http.ResponseWriter, r *http.Request) {
		//if r.Method == http.MethodPost {
		users = append(users, user{
			Name: fmt.Sprintf("%d", len(users)+1),
		})
		//	}
	})
	http.HandleFunc("/users", func(rw http.ResponseWriter, r *http.Request) {
		//if r.Method == http.MethodPost {
		response, err := json.Marshal(users)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(string(response))
		rw.Write(response)
		//	}
	})
	http.HandleFunc("/users/", func(rw http.ResponseWriter, r *http.Request) {
		//if r.Method == http.MethodPost {
		fmt.Println(r.URL.Path)
		userID := strings.TrimPrefix(r.URL.Path, "/users/")
		id, err := strconv.Atoi(userID)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		response, err := json.Marshal(users[id-1])
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.Write(response)
		//	}
	})
	http.ListenAndServe(":8080", nil)
}
