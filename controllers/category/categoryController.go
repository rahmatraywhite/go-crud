package categoryController

import (
	"go-web-native/entities"
	categoryModel "go-web-native/models/category"
	"html/template"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categoryModel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	template, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}

	template.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		template, err := template.ParseFiles("views/category/create.html")
		if err != nil {
			panic(err)
		}

		template.Execute(w, nil)
	}

	if r.Method == "POST" {
		var category entities.Category

		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpddatedAt = time.Now()

		if ok := categoryModel.Create(category); !ok {
			template, _ := template.ParseFiles("views/category/create.html")
			template.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}

}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
