package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type PagePromo struct {
	Classe          string
	Filliere        string
	Niveau          string
	NombreEtudiants int
	Etudiants       []Etudiant
}

type Etudiant struct {
	Nom    string
	Prenom string
	Age    int
	Sexe   bool //true := homme
}

func main() {
	temp, err := template.ParseGlob("./challenge1/template/*.html")

	if err != nil {
		fmt.Println(fmt.Sprint("erreur %s", err.Error()))
		return
	}
	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		ListeEtudiants := []Etudiant{{"gourrin", "dimitri", 17, true}, {"Xerli", "Chinois", 152, false}}

		PageData := PagePromo{"B1 Informatique", "Informatique", " Bachelor 1", 45, ListeEtudiants}
		temp.ExecuteTemplate(w, "Promo", PageData)

	})
	fileServer := http.FileServer(http.Dir("./challenge1/assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	fileServerStyles := http.FileServer(http.Dir("./challenge1/styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", fileServerStyles))
	http.ListenAndServe("localhost:8000", nil)
}
