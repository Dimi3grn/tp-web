package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type PageChange struct {
	Vues      int
	Condition bool
}

var PageData PageChange

func main() {
	temp, err := template.ParseGlob("./challenge3/template/*.html")

	if err != nil {
		fmt.Println(fmt.Sprint("erreur %s", err.Error()))
		return
	}
	http.HandleFunc("/user/form", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "Form", nil)
	})

	http.HandleFunc("/user/treatment", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "Treatment", nil)
	})

	fileServer := http.FileServer(http.Dir("./challenge3/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.ListenAndServe("localhost:8000", nil)
}
