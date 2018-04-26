package main

import (
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
)

const (
	HTML_TEMPLATE = `<html>
    <head>
	    <title>Hello World</title>
	    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
	</head>
	<body>
		<div class="container">
			<div class="jumbotron">
				<h1>Hello {{.Name}}!</h1>
			</div>
        </div>
		<!-- jQuery library -->
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
		<!-- Latest compiled JavaScript -->
		<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
	</body>
</html>
`
)

type Person struct {
	Name string
}

func prettyPrintRequest(req *http.Request) {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(requestDump))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	prettyPrintRequest(r)
	t := template.New("person template")

	t, _ = t.Parse(HTML_TEMPLATE)

	personName := r.URL.Query().Get("name")
	if personName == "" {
		personName = "World"
	}

	p := Person{Name: personName}
	t.Execute(w, p)
}

func main() {
	port := ":8081"
	http.HandleFunc("/", handleRequest)

	log.Printf("Listening on port [%s]", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
