package main

import (
    "net/http"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.HandleFunc("/projects", projectsHandler)
    http.HandleFunc("/contact", contactHandler)

    http.ListenAndServe(":8080", nil)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "templates/projects.html")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "templates/contact.html")
}

