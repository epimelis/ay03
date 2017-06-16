package main

import ( "fmt"
	"net/http"
	"errors"
	"strings"
	"html/template"
	"github.com/epimelis/ay03/data"
)

func p(a ...interface{}) {
	fmt.Println(a)
}

func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
	url :=[]string {"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error){
	cookie, err :=request.Cookie("_cookie")
	if err ==nil {
		sess =data.Session{Uuid: cookie.Value}
		if ok, _ :=sess.Check(); !ok{
			err = errors.New("Invalid session")
		}
	}
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

