package stom

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ResponseWriter struct {
	http.ResponseWriter
	Status      int
	StartTime   time.Time
	wroteHeader bool
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(data)
}

func (w *ResponseWriter) WriteHeader(status int) {
	w.wroteHeader = true
	w.Status = status
	w.ResponseWriter.WriteHeader(status)
}

// WriteString is a Printf-like helper function that will write a text/plain header
// and string to the ResponseWriter.
func WriteString(w http.ResponseWriter, format string, args ...interface{}) {
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, format, args...)
}

// WriteJSON is a helper function that will write the JSON encoded version of v
// to the ResponseWriter and set the application/json Content-Type header.
func WriteJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Add("Content-Type", "application/json")
	data, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
