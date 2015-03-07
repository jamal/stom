package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jamal/stom"
)

type Logger struct {
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sw, ok := w.(*stom.ResponseWriter)
	if !ok {
		return
	}

	// RemoteAddr Time "Method Path HTTP version" Status Time
	fmt.Printf("%s [%s] \"%s %s %s\" %d %.10fs\n",
		r.RemoteAddr,
		time.Now().Format(time.RFC1123Z),
		r.Method,
		r.URL.Path,
		r.Proto,
		sw.Status,
		time.Since(sw.StartTime).Seconds(),
	)
}
