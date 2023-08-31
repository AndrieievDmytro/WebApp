package render

import (
	"WebApp/helper"
	"net/http"
	"text/template"
)

func RenderTempalte(w http.ResponseWriter, tmpl string) {
	parsedTempalte, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTempalte.Execute(w, nil)
	helper.CheckErr(err)
}
