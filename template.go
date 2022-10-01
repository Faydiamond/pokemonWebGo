package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/**/*.html"))
var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

/* Page to render*/
func Index(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Pokemones")
	renderTemplate(rw, "index.html", nil)
}

func History(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Hikstory")
	renderTemplate(rw, "history.html", nil)
}

func Seasons(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, " temporadas ")
	renderTemplate(rw, "seasons.html", nil)
}

func Contact(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Contacto")
	renderTemplate(rw, "contact.html", nil)
}

/* End page to render */

/* Render templates */
func renderTemplate(rw http.ResponseWriter, file string, data interface{}) {
	err := templates.ExecuteTemplate(rw, file, data)
	if err != nil {
		//http.Error(rw, "No fue posible renderizar el template", http.StatusInternalServerError)
		errTemplate(rw, http.StatusInternalServerError)
	}
}

/* end render templates */

/*error template*/
func errTemplate(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	errorTemplate.Execute(rw, nil)
}

/*end error template*/

func main() {
	staticFile := http.FileServer(http.Dir("static"))
	//mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/historia", History)
	mux.HandleFunc("/temporadas", Seasons)
	mux.HandleFunc("/contacto", Contact)
	mux.Handle("/static/", http.StripPrefix("/static/", staticFile))
	//mensajes por consola
	fmt.Println("Se esta ejecutando el servidor en el puerto 9000")
	fmt.Println("Servidor corriendo  http://localhost:9000/")

	//Servidor
	server := &http.Server{
		Addr:    "localhost:9000",
		Handler: mux,
	}

	server.ListenAndServe()
}
