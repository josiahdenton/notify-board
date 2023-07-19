package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
    "github.com/josiahdenton/notify-board/internal/logging"
)

func root(rw http.ResponseWriter, req *http.Request) {
    infoLogger := logging.New(logging.INFO)
    errLogger := logging.New(logging.ERROR)

    infoLogger.Println("Request received for '/'")

    templ, err := template.ParseFiles("web/templates/index.html")
    if err != nil {
        fmt.Fprintf(rw, "ERROR!")
    }
    infoLogger.Println("Rendering root template")
    err = templ.Execute(rw, nil)
    if err != nil {
        errLogger.Printf("Failed to parse root template %v", err)
    }
}


func click(rw http.ResponseWriter, req *http.Request) {
    infoLogger := logging.New(logging.INFO)
    // errLogger := logging.NewLogger(logging.ERROR)
    infoLogger.Println("Request received for '/click'")
    fmt.Fprintf(rw, "<div>Hello World!</div>")
}


func main() {
    // TODO - handle interrup SIG when we have DB conn

    fmt.Println("Starting notify-board...")
    http.HandleFunc("/", root)
    http.HandleFunc("/click", click)

    err := http.ListenAndServe(":8001", nil)
    if errors.Is(err, http.ErrServerClosed) {
        fmt.Println("no problemo")
    } else {
        log.Fatalln("OH NOOOO", err) 
    }
}
