package main

import (
	"fmt"
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Username string
	Password []byte
	First string
	Last string
}

var tpl *template.Template
var dbUsers = map[string]user{};
var dbSessions = map[string]string{};


func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	tpl.ExecuteTemplate(w, "index.html", u)
}

func welcome(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := getUser(req)
	tpl.ExecuteTemplate(w, "welcome.html", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/welcome", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		// check if username already exists
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		sid := uuid.NewV4()
		c := &http.Cookie{
			Name: "session",
			Value: sid.String(),
		}

		http.SetCookie(w, c)

		dbSessions[c.Value] = un

		bp, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)

		u := user{
			un, bp, f, l,
		}

		dbUsers[un] = u

		http.Redirect(w, req, "/welcome", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.html", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/welcome", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodPost {
		var u user
		un := req.FormValue("username")
		p := req.FormValue("password")
		
		// check if username exists
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "username not found", http.StatusForbidden)
			return
		}
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "password doesn't match", http.StatusForbidden)
			return
		}

		// everything okay so create cookie
		sid := uuid.NewV4()
		c := &http.Cookie{
			Name: "session",
			Value: sid.String(),
		}

		http.SetCookie(w, c)
		dbSessions[c.Value] = un
		http.Redirect(w, req, "/welcome", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	delete(dbSessions, c.Value)
	c = &http.Cookie{
		Name: "session",
		Value: "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		fmt.Println("l")
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}

// gets user if the user exists
func getUser(req *http.Request) user {
	var u user
	c, err := req.Cookie("session")
	if err != nil {
		// return empty user
		return u
	}
	// if the user is already registered
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}