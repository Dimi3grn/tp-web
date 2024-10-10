package main

import (
	"fmt"
	"net/http"
	"regexp"
	"text/template"
)

type PageInfo struct {
	nom    string
	prenom string
	bday   string
	sexe   string
}

var PageData PageInfo

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

		PageData.nom = r.FormValue("nom")
		PageData.prenom = r.FormValue("prenom")
		PageData.bday = r.FormValue("bday")
		PageData.sexe = r.FormValue("sexe")
		http.Redirect(w, r, "/user/display", http.StatusSeeOther)
	})
	type displayPage struct {
		Name   string
		Prenom string
		Bday   string
		Sexe   string
	}

	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) {

		displayData := displayPage{PageData.nom, PageData.prenom, PageData.bday, PageData.sexe}

		temp.ExecuteTemplate(w, "Display", displayData)
	})

	http.HandleFunc("/erreur", func(w http.ResponseWriter, r *http.Request) {
		code, message := r.FormValue("code"), r.FormValue("message")
		fmt.Fprintf(w, "Erreur %s - %s", code, message)
	})

	fileServer := http.FileServer(http.Dir("./challenge3/styles"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.ListenAndServe("localhost:8000", nil)
}
