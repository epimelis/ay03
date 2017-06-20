package main

import ( "fmt"
	"net/http"
	"errors"
	"strings"
	"html/template"
	"github.com/epimelis/ay03/data"
	"log"
)

var logger *log.Logger
func p(a ...interface{}) {
	fmt.Println(a)
}

func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
	url :=[]string {"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error){
	fmt.Println("utils.session_#1")
	cookie, err :=request.Cookie("_cookie")
	fmt.Println("utils.session_#2")
	if err ==nil {
		fmt.Println("utils.session_#2.1")
		sess =data.Session{Uuid: cookie.Value}
		fmt.Println("utils.session_#2.2 : ", sess.Uuid)
		if ok, _ :=sess.Check(); !ok{
			err = errors.New("Invalid session")
			fmt.Println("utils.session_ERR1!! : ", err)
		} else {
			fmt.Println("utils.session_#2.3 - OK. Session is valid")
		}
	} else {
		fmt.Println("utils.session_ERR2 - request.Cookie failed!!")
	}
	fmt.Println("utils.session_#3")
	return
}

func parseTemplateFiles(filenames ...string) (t *template.Template) {

	var files []string
	t=template.New("layout")
	for _, file :=range filenames{
		files=append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))

	return
}

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file :=range filenames {
		files =append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates :=template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

// for logging
func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

// version
func version() string {
	return "1.0"
}


