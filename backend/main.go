package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	css := http.FileServer(http.Dir("../client/style"))
	http.Handle("/style/", http.StripPrefix("/style/", css))

	js := http.FileServer(http.Dir("../client/scripts"))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	resources := http.FileServer(http.Dir("../client/resources"))
	http.Handle("/resources/", http.StripPrefix("/resources/", resources))

	tmpl, err := template.ParseGlob("../client/pages/*.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tmpl.DefinedTemplates())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		if r.URL.Path == "/" {
			err = tmpl.ExecuteTemplate(w, "index.html", nil)
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	log.Println("Server started")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
