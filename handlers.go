package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTempalte(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	RenderTempalte(w, "about.page.tmpl")
}
