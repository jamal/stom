# Stom

Stom is a lightweight library that helps you write web applications and APIs in [Go](http://golang.org/). The goal of
stom is to be as non-intrusive as possible, by just adding useful features that are commonly used in web applications
without adding any magic. It tries to stick to the standard http library as much as possible, by using the standard http
handlers (`func(http.ResponseWriter, *http.Request)`) without adding any extra fluff. It also recommends some usage
patterns based in regards to handling application Context.

For routing, it uses [github.com/julienschmidt/httprouter](http://github.com/julienschmidt/httprouter), and exposes
route variables through the standard `http.Request.FormValue`.

Please note that this project is still extremely young and may change often. The API should remain stable, though.

## Usage

```go
package main

import (
	"log"
	"net/http"

	"github.com/jamal/stom"
	"github.com/jamal/stom/middleware"
)

func Index(w http.ResponseWriter, r *http.Request) {
	stom.WriteString(w, "Welcome")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	storm.WriteString(w, "Hello, %s", r.FormValue("name"))
}

func main() {
	s := stom.New()
	s.UseAfter(middleware.Logger{})
	s.Get("/", Index)
	s.Get("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe(":8080", s))
}
```
See [example/main.go](example/main.go) for a more detailed example.

## Request Context

Request context is not handled by this library, because it isn't needed. Instead, it offers the [`Context`](http://godoc.org/github.com/jamal/stom#Context) interface as an suggestion on how to approach this problem. With this pattern, you only need to fetch the request context on handlers that need it, and `NewContext` should handle any parsing of request parameters that are needed (such as reading a session cookie and fetching the User object). For example:

```go
type Context {
	User *User
}

func NewContext(r *http.Request) *Context {
	ctx := new(Context)
	sessionID, _ := r.Cookie("session_id")
	user, _ := FetchSessionUser(sessionID)
	ctx.User = user
	return ctx
}

func ContextHandler(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(r)
	
}
```
