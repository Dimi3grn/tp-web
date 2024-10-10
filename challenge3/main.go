package main

import (
	"fmt"
	"net/http"
	"regexp"
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

type PageChange struct {
	Vues      int
	Condition bool
}

var ChangeData PageChange

type PageInfo struct {
	nom    string
	prenom string
	bday   string
	sexe   string
}

var FormData PageInfo

func main() {
	temp, err := template.ParseGlob("./challenge3/template/*.html")

	if err != nil {
		fmt.Println(fmt.Sprint("erreur %s", err.Error()))
		return
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		ListeEtudiants := []Etudiant{{"Gourrin", "Dimitri", 17, true}, {"Xerli", "Chinois", 152, false}}

		PageData := PagePromo{"B1 Informatique", "Informatique", " Bachelor 1", len(ListeEtudiants), ListeEtudiants}
		temp.ExecuteTemplate(w, "Promo", PageData)

	})

	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		ChangeData.Vues++
		if ChangeData.Vues%2 == 0 {
			ChangeData.Condition = false
		} else {
			ChangeData.Condition = true
		}
		temp.ExecuteTemplate(w, "Change", ChangeData)
	})

	http.HandleFunc("/user/form", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "Form", nil)
	})

	http.HandleFunc("/user/treatment", func(w http.ResponseWriter, r *http.Request) {

		var checkValue bool
		checkValue, _ = regexp.MatchString(`^[A-Za-z]{1,32}$`, r.FormValue("nom"))
		if !checkValue {
			http.Redirect(w, r, "/erreur?code=400&message=Oups les données sont invalides", http.StatusMovedPermanently)
			return
		}
		checkValue, _ = regexp.MatchString(`^[A-Za-z]{1,32}$`, r.FormValue("prenom"))
		if !checkValue {
			http.Redirect(w, r, "/erreur?code=400&message=Oups les données sont invalides", http.StatusMovedPermanently)
			return
		}
		checkValue, _ = regexp.MatchString(`^(masculin|féminin|autre)$`, r.FormValue("sexe"))
		if !checkValue {
			http.Redirect(w, r, "/erreur?code=400&message=Oups les données sont invalides", http.StatusMovedPermanently)
			return
		}

		FormData.nom = r.FormValue("nom")
		FormData.prenom = r.FormValue("prenom")
		FormData.bday = r.FormValue("bday")
		FormData.sexe = r.FormValue("sexe")
		http.Redirect(w, r, "/user/display", http.StatusSeeOther)
	})
	type displayPage struct { //sert a rien mais azy je me suis cassé le crane pour rien
		Name   string
		Prenom string
		Bday   string
		Sexe   string
	}

	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) {

		displayData := displayPage{FormData.nom, FormData.prenom, FormData.bday, FormData.sexe}

		temp.ExecuteTemplate(w, "Display", displayData)
	})

	http.HandleFunc("/erreur", func(w http.ResponseWriter, r *http.Request) {
		code, message := r.FormValue("code"), r.FormValue("message")
		fmt.Fprintf(w, "Erreur %s - %s", code, message)
	})

	fileServer := http.FileServer(http.Dir("./challenge3/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.ListenAndServe("localhost:8000", nil)
}
