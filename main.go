package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

type result struct {
	Url    string
	Time   string
	Status bool
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		str := strings.Replace(string(r.URL.Path), "/", "", 1)
		res, _ := json.Marshal(CheckServer(str))
		fmt.Fprintf(w, string(res))
	})
	http.ListenAndServe("127.0.0.1:1234", nil)
}

func CheckServer(uri string) result {
	var res result
	if strings.Count(uri, ":") == 0 {
		uri += ":80"
	}
	res.Url = uri
	if strings.Count(uri, ".") == 0 {
		res.Status = false
		return res
	}
	timeout := time.Duration(5 * time.Second)
	t1 := time.Now()
	_, err := net.DialTimeout("tcp", uri, timeout)
	res.Time = time.Now().Sub(t1).String()
	if err != nil {
		res.Status = false
	} else {
		res.Status = true
	}
	return res
}
