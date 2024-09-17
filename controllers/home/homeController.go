package homeController

import (
	"html/template"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}

	template.Execute(w, nil)
}
