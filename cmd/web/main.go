package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.LUTC)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)

	// Initialize a instance of app
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Use the http.NewServerMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	fileServer := http.FileServer(http.Dir("./ui/static"))
	// mux.Handle("/static/", fileServer)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Use the http.ListenAndServe() function to start a new web server
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
