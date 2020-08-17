package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func s(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Any Code")
}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Dog")
}

func m(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(res, "something.gohtml", "Louis")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {
	http.HandleFunc("/", s)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", m)

	http.ListenAndServe(":8080", nil)
}
