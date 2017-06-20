package main

import (
	"net/http"
	"github.com/epimelis/ay03/data"
	"fmt"
)

// GET /threads/new
// Show the new thread form page
func newThread(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /signup
// Create the user account
func createThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			danger(err, "Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func readThread(writer http.ResponseWriter, request *http.Request) {
	p("readThread_#1")
	vals := request.URL.Query()
	p("readThread_#2")
	uuid := vals.Get("id")
	p("readThread_#3 : ", uuid)

	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		p("readThread_ERR1 !!!", err)
		error_message(writer, request, "Cannot read thread")
	} else {
		p("readThread_#4")
		_, err := session(writer, request)
		p("readThread_#5")

		if err != nil {
			p("readThread_#5.1")

			generateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
		} else {
			p("readThread_#5.2")

			generateHTML(writer, &thread, "layout", "private.navbar", "private.thread")
		}
	}
	p("readThread_#6")

}

// POST /thread/post
// Create the post

func postThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			error_message(writer, request, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
