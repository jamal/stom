package stom

import (
	"net/http"
	"net/url"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Handle func(http.ResponseWriter, *http.Request)

type Server struct {
	PanicHandler func(http.ResponseWriter, *http.Request, interface{})

	router    *httprouter.Router
	useBefore []http.Handler
	useAfter  []http.Handler
}

func New() *Server {
	s := new(Server)
	s.router = httprouter.New()

	// Some default handlers
	s.router.PanicHandler = s.panicHandler
	return s
}

func (s *Server) ServeHTTP(hw http.ResponseWriter, r *http.Request) {
	// We use a custom ResponseWriter to capture the response Status
	w := new(ResponseWriter)
	w.ResponseWriter = hw
	w.StartTime = time.Now()

	for _, m := range s.useBefore {
		m.ServeHTTP(w, r)
	}

	s.router.ServeHTTP(w, r)

	for _, m := range s.useAfter {
		m.ServeHTTP(w, r)
	}
}

func (s *Server) Get(path string, handle Handle) {
	s.Handle("GET", path, handle)
}

func (s *Server) Head(path string, handle Handle) {
	s.Handle("HEAD", path, handle)
}

func (s *Server) Post(path string, handle Handle) {
	s.Handle("POST", path, handle)
}

func (s *Server) Put(path string, handle Handle) {
	s.Handle("PUT", path, handle)
}

func (s *Server) Delete(path string, handle Handle) {
	s.Handle("DELETE", path, handle)
}

func (s *Server) Handle(method, path string, handle Handle) {
	s.router.Handle(method, path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Use http.Request.FormValue to get route-defined URL parameters
		if r.Form == nil {
			r.Form = make(url.Values)
		}
		for _, p := range ps {
			r.Form.Add(p.Key, p.Value)
		}
		handle(w, r)
	})
}

func (s *Server) Use(middleware http.Handler) {
	s.useBefore = append(s.useBefore, middleware)
}

func (s *Server) UseAfter(middleware http.Handler) {
	s.useAfter = append(s.useAfter, middleware)
}

func (s *Server) panicHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if s.PanicHandler != nil {
		s.PanicHandler(w, r, err)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		WriteString(w, "500 internal server error")
	}
}
