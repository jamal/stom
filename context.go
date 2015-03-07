package stom

import "net/http"

// The Context interface allows you to implement a shared context that is
// needed among a group of request types. NewContext should process the
// Request for any information that is necessary for a particular request
// type. For example, an endpoint that requires authentication could parse
// the authentication header, or a session cookie, and set User in the
// context. The idea is that Context is not generated by stom because it
// may not be needed in every request. In your request handler simply call
// NewContext(r) when you need it.
//
// A simple example:
//
//  type User struct {
//      ID int
//      Name string
//  }
//
//  type Context struct {
//      User *User
//  }
//
//  func NewContext(r *http.Request) *Context {
//      ctx := new(Context)
//      // Error handling omitted for simplicity
//      sessionID, _ := r.Cookie("session_id")
//      user, _ := FetchSessionUser(sessionID)
//      ctx.User = user
//      return ctx
//  }
//
//  func HelloName(w http.ResponseWriter, r *http.Request) {
//      ctx := NewContext(r)
//      if ctx.User != nil {
//          fmt.Fprintf(w, "Hello, %s", ctx.User.Name)
//      } else {
//          fmt.Fprintf(w, "Hello, I don't know who you are!")
//      }
//  }
//
type Context interface {
	NewContext(r *http.Request) Context
}
