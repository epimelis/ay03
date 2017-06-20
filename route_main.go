package main

import (
	"net/http"
	"fmt"
	"github.com/epimelis/ay03/data"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}


func index(writer http.ResponseWriter, request *http.Request) {
	p("index_#1")
	threads, err :=data.Threads()
	p("index_#2")
	if err !=nil {
		p("index_ERR1 - Cannot get threads!!")
		error_message(writer, request, "Cannot get threads")
	} else {
		p("index_#2.1")
		_, err := session(writer, request)
		p("index_#2.2")
		if err != nil {
			p("index_#2.2.1 - No session found")
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			p("index_#2.2.2 - Has valid session")
			generateHTML(writer, threads, "layout", "private.navbar", "index")
			fmt.Println("ii3")
		}
		p("index_#3")
	}
}
