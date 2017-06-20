package main

import (
	"net/http"
	"github.com/epimelis/ay03/data"
	"fmt"
)


func login(writer http.ResponseWriter, request *http.Request) {
	p("login()_#1")
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
	p("login()_#2")
}


func signup(writer http.ResponseWriter, request *http.Request) {
	p("signup()_#1")
	generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
	p("signup()_#1")
}

func signupAccount(writer http.ResponseWriter, request *http.Request) {
	err :=request.ParseForm()
	p("signupAccount_#1")
	if err !=nil {
		p("signupAccount_ERR1 !! - Cannot parse form")
		danger(err)
	}
	user := data.User{
		Name: request.PostFormValue("name"),
		Email: request.PostFormValue("email"),
		Password: request.PostFormValue("password"),

	}
	if err := user.Create(); err !=nil {
		danger(err, "Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}

func authenticate(writer http.ResponseWriter, request *http.Request) {
	err :=request.ParseForm()
	p("authenticate_#1")
	if err !=nil {
		p("authenticate_ERR1 !! - Cannot parse form")
		danger(err, "Cannot parse form")
	}
	p("authenticate_#2")
	fmt.Println("2a", request.PostFormValue("email"))
	fmt.Println("2b", request.PostFormValue("password"))

	user, err :=data.UserByEmail(request.PostFormValue("email"))
	if err !=nil {
		danger(err, "Cannot find user")
	}
	p("authenticate_#3a", user.Id)
	p("authenticate_#3b", user.Password)
	p("authenticate_#3C", user.Email)

	if user.Password==request.PostFormValue("password") {
		p("authenticate_#3.1")
		session, err :=user.CreateSession()
		p("authenticate_#3.1a : ", session.Uuid, "###", session.Email, "@@@", session.UserId)
		if err !=nil {
			p("authenticate_#3.1_ERR - cannot create session")
			danger(err, "cannot create session")
		}
		p("authenticate_#3.2")
		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.Uuid,
			HttpOnly: true,
		}
		p("authenticate_#3.3")
		http.SetCookie(writer,&cookie)
		p("authenticate_#3.4")
		http.Redirect(writer, request, "/", 302)
	} else {
		p("authenticate_#3.5")
		http.Redirect(writer, request, "/login", 302)
	}
	p("authenticate_#4")
}

func logout(writer http.ResponseWriter, request *http.Request) {
	p("logout_#1")
	cookie, err := request.Cookie("_cookie")
	p("logout_#2")
	if err !=http.ErrNoCookie {
		p("logout_#2a - cookie is present - ", "@@@", err, "###", http.ErrNoCookie)
		//warning(err, "Failed to get cookie")
		p("logout_#2b")
		session := data.Session{Uuid: cookie.Value}
		p("logout_#2c")
		session.DeleteByUUID()
	} else {
		p("logout_ERR1 - cookie is not present!!")

	}
	http.Redirect(writer, request, "/", 302)
}

