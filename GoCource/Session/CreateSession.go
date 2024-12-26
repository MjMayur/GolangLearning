package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", createSession)
	http.HandleFunc("/bar", handleBar)
	http.ListenAndServe(":8080", nil)
}

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func createSession(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		token := uuid.New()
		c = &http.Cookie{
			Name:  "session",
			Value: token.String(),
		}
		http.SetCookie(res, c)
	}
	// if the user exists already, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func handleBar(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}
