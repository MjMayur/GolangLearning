package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName    string
	Email       string
	Password    []byte
	ConfirmPass string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func signup(res http.ResponseWriter, req *http.Request) {
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

	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		e := req.FormValue("email")
		p, err := bcrypt.GenerateFromPassword([]byte(req.FormValue("password")), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal Server error", http.StatusInternalServerError)
			return
		}
		cp := req.FormValue("confirmpass")
		u = user{un, e, p, cp}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func handleBar(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")

	if err != nil {

		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		res.Write([]byte(`
            <script>
                alert("Please sign up first");
                window.location.href = "/";
            </script>
        `))
		return
	}

	un, ok := dbSessions[c.Value]
	if !ok {

		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		res.Write([]byte(`
            <script>
                alert("Please sign up first***************");
                window.location.href = "/";
            </script>
        `))
		return
	}

	// If session is valid, get user and render template
	u := dbUsers[un]
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}

func logout(res http.ResponseWriter, req *http.Request) {
	c, _ := req.Cookie("session")

	c = &http.Cookie{
		Name:   "session", // Replace "session" with the name of the cookie you want to clear
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, c)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
