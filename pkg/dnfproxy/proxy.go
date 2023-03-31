package dnfproxy

import (
	"fmt"
	"html"
	"net/http"
	"time"
)

type Proxy interface {
	Start() error
}

type DnfProxy struct {
	port   uint
	server *http.Server
}

func New(port uint) Proxy {
	return DnfProxy{port: port}
}

func (d DnfProxy) mainHandler() {

}
func (d DnfProxy) Start() error {
	d.server = &http.Server{
		Addr:           fmt.Sprintf(":%d", d.port),
		Handler:        http.DefaultServeMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	return d.server.ListenAndServe())
}
