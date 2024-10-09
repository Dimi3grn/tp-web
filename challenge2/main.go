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
	temp, err := template.ParseGlob("./challenge2/template/*.html")

	if err != nil {
		fmt.Println(fmt.Sprint("erreur %s", err.Error()))
		return
	}
	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		PageData.Vues++
		if PageData.Vues%2 == 0 {
			PageData.Condition = false
		} else {
			PageData.Condition = true
		}
		temp.ExecuteTemplate(w, "Change", PageData)
	})
	fileServer := http.FileServer(http.Dir("./challenge2/styles"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.ListenAndServe("localhost:8000", nil)
}
