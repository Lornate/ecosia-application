package controller

import (
	"ecosia-application/model"
	"html/template"
	"net/http"
	"os"
)

const (
	error = "Please tell me your favorite tree"
)

func GetFavouriteTree(w http.ResponseWriter, r *http.Request) {

	values, ok := r.URL.Query()["favoriteTree"]

	if !ok || len(values[0]) < 1 {

		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, error, http.StatusBadRequest)
		return
	} else {

		tree := model.FavouriteTree{Tree: values[0]}
		wd, err := os.Getwd()
		tmpl, err := template.ParseFiles(wd + "/templates/base.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, tree); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		return
	}

}
