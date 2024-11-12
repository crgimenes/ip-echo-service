package main

import (
	"embed"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"text/template"
	"time"
)

//go:embed assets
var assets embed.FS

func getParameters(prefix string, r *http.Request) []string {
	path := strings.TrimPrefix(r.URL.Path, prefix)
	path = strings.TrimSuffix(path, "/")
	path = strings.TrimSpace(path)
	a := strings.Split(path, "/")

	b := make([]string, len(a))
	i := 0

	for _, v := range a {
		if v != "" {
			b[i] = v
			i++
		}
	}

	return b[:i]
}

func handlerMain(w http.ResponseWriter, r *http.Request) {
	parameters := getParameters("/", r)

	lang := r.Header.Get("Accept-Language")
	log.Printf("Accept-Language: %s", lang)

	mode := "html"
	if len(parameters) > 0 {
		mode = parameters[0]
	}

	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
		log.Printf("RemoteAddr: %v", ip)
		h, _, err := net.SplitHostPort(ip)
		if err != nil {
			log.Printf("Error: %+v", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "internal server error")
			return
		}

		ip = h
	}

	switch mode {
	case "json":
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"ip": "%s"}`, ip)
	case "text":
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%s", ip)
	default:
		w.Header().Set("Content-Type", "text/html")
		tpl, err := template.ParseFS(assets, "assets/index.html")
		if err != nil {
			log.Printf("Error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			fmt.Fprintf(w, "Error: %s", err)
			return
		}
		type data struct {
			IP string
		}

		d := data{IP: ip}
		tpl.Execute(w, d)
		return
	}
}

func main() {

	port := flag.String("port", "8001", "Port to listen on")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handlerMain))

	s := &http.Server{
		Handler:        mux,
		Addr:           fmt.Sprintf(":%s", *port),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Starting server on port %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
