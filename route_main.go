package main

import (
	"net/http"
	"fmt"
	"github.com/epimelis/ay03/data"
)

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err :=data.Threads()
	if err !=nil {
		error_message(writer, request, "Cannot get threads")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(writer, threads, "layout", "private.navbar", "index")
			fmt.Println("ii3")
		}
	}
}
