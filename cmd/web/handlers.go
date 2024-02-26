package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    app.notFound(w)
    return
  }

  // Initialize a slice containing the paths to the two files. Its important 
  // to note that the file containg our base template must be the "first"
  files := []string{
    "./ui/html/base.tmpl",
    "./ui/html/partials/nav.tmpl",
    "./ui/html/pages/home.tmpl",
  }

  // Use the template.ParseFiles() funtion to read the files and store the 
  // templates in a template set. Notice that we use ... to pass the contants
  // of the file slice as variadic arguments.
  ts, err := template.ParseFiles(files...)
  if err != nil {
    app.serverError(w, r, err)
    return
  }

  // Use the ExecuteTemplate() method to write the content to the "base"
  // template as the response body.
  err = ts.ExecuteTemplate(w, "base", nil)
  if err != nil {
    app.serverError(w, r, err)
  }
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    app.notFound(w)
    return
  }

  fmt.Fprintf(w, "Dipslay a specific snipper with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    w.Header().Set("Allow", http.MethodPost)
    app.clientError(w, http.StatusMethodNotAllowed)
    return
  }

  w.Write([]byte("Create a new snippet"))
}
