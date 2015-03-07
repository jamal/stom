package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/jamal/stom"
	"github.com/jamal/stom/middleware"
)

var db *sql.DB

type Context struct {
	DB   *sql.DB
	User *User
}

// User is used as an example of something you could store in Context
type User struct {
	ID   int
	Name string
}

func NewContext(r *http.Request) *Context {
	return &Context{db, &User{1, "Jamal"}}
}

func Index(w http.ResponseWriter, r *http.Request) {
	stom.WriteString(w, "Welcome")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(r)
	stom.WriteString(w, "Hello, %s", ctx.User.Name)
}

func HelloName(w http.ResponseWriter, r *http.Request) {
	stom.WriteString(w, "Hello, %s", r.FormValue("name"))
}

func main() {
	s := stom.New()
	s.UseAfter(middleware.Logger{})

	s.Get("/", Index)
	s.Get("/hello", Hello)
	s.Get("/hello/:name", HelloName)

	log.Fatal(http.ListenAndServe(":8080", s))
}
