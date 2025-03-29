/*
Exercises

1.	Write a small web server that returns the current time in RFC 3339 format
	when you send it a GET command. You can use a third-party module if you’d
	like.

2.	Write a small middleware component that uses JSON structured logging to
	log the IP address of each incoming request to your web server.

3.	Add the ability to return the time as JSON. Use the Accept header to
	control whether JSON or text is returned (default to text). The JSON
	should be structured as follows:
{
    "day_of_week": "Monday",
    "day_of_month": 10,
    "month": "April",
    "year": 2023,
    "hour": 20,
    "minute": 15,
    "second": 20
}
*/

package main

import (
	"net/http"
	"time"
)

type TimeHandler struct{}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().Format(time.RFC3339)
		w.Write([]byte(t + "\n"))
	})

	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
