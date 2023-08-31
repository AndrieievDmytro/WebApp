package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var (
	// templateCache  = make(map[string]*template.Template)
	pageTemplates  = "./templates/*.page.tmpl"
	layoutTemlates = "./templates/*.layout.tmpl"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal("error parsing templates:", err)
	}

	//get requested template
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	// Check if something wrong with templates stored in the map, do not execute directly.
	_ = t.Execute(buf, nil)

	//render the tempate
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all the files end with *.page.tmpl from ./templates

	pages, err := filepath.Glob(pageTemplates)
	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		//populate template set
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob(layoutTemlates)
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(layoutTemlates)
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	//check if we already have template in our cache
// 	_, inMap := templateCache[t]
// 	if !inMap {
// 		log.Println("creating template and adding to cache")
// 		//create the template
// 		err = CreateTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		log.Printf("%s template created", t)
// 	} else {
// 		log.Printf("%s template already in cache", t)
// 	}

// 	tmpl = templateCache[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func CreateTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	//parse the templates
// 	tmpl, err := template.ParseFiles(templates...)

// 	if err != nil {
// 		return err
// 	}
// 	//add templates to cache (map)
// 	templateCache[t] = tmpl
// 	return nil
// }
