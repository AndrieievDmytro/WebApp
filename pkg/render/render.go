package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTempalte(w http.ResponseWriter, tmpl string) {
	parsedTempalte, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTempalte.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing tempalte")
	}
}
