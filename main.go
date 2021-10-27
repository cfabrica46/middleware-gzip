package main

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"

	"github.com/cfabrica46/middleware-gzip/gingzip"
)

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func main() {
	go gingzip.GinRouter()

	handler := http.HandlerFunc(index)
	http.Handle("/", middleware(handler))
	log.Println("ListenAndServe on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		gzr.ResponseWriter.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(gzr, r)
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("index"))
	if err != nil {
		log.Println(err)
		return
	}
}
