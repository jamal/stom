# Stom

Stom is a lightweight library that helps you write web applications and APIs in [Go](http://golang.org/). The goal of
stom is to be as non-intrusive as possible, by just adding useful features that are commonly used in web applications
without adding any magic. It tries to stick to the standard http library as much as possible, by using the standard http
handlers (`func(http.ResponseWriter, *http.Request)`) without adding any extra fluff. It also recommends some usage
patterns based in regards to handling application Context.

For routing, it uses [github.com/julienschmidt/httprouter](http://github.com/julienschmidt/httprouter), and exposes
route variables through the standard `http.Request.FormValue`.

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

func main() {
	s := stom.New()
	s.UseAfter(middleware.Logger{})
	s.Get("/", Index)
	log.Fatal(http.ListenAndServe(":8080", s))
}
```
