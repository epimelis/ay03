package main

import(
	"net/http")


func main() {

	mux :=http.NewServeMux()
	filesHandler :=http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", filesHandler))


	mux.HandleFunc("/", index)
	/*
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/signup", signup)
	*/
	server :=&http.Server{
		Addr: "0.0.0.0:7777",
		Handler: mux,
	}
	server.ListenAndServe()
}