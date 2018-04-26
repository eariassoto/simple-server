package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Person struct {
	Name string
}

func main() {
	port := ":8081"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("person template")

		templateFile, err := ioutil.ReadFile("index.html")
		if err != nil {
			fmt.Fprintf(w, "Error from server")
		}
		s := string(templateFile)

		t, _ = t.Parse(s)

		personName := r.URL.Query().Get("name")
		if personName == "" {
			personName = "World"
		}

		p := Person{Name: personName}
		t.Execute(w, p)
	})

	log.Printf("Listening on port [%s]", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
