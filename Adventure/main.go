package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
type Story struct {
	Title   string   `json:"title"`
	Body    []string `json:"story"`
	Options []Option `json:"options"`
}

func main() {

	//1Read JSON and convert to map
	jsonStr, _ := os.ReadFile("gopher.json")

	var jsonMap map[string]Story

	_ = json.Unmarshal(jsonStr, &jsonMap)

	//2Create Template engine
	var tmplFile = `

<head>
	<meta charset="utf-8">
	<title>Choose Your Own Adventure</title>
</head>
<body>
<h1>{{.Title}}</h1>
{{range .Body}}
	<p>{{.}}</p>
{{end}}
{{range .Options}}
<a href="/{{.Arc}}">{{.Text}}</a></br>
{{end}}
</body>
`
	tmp, err := template.New("webTmp").Parse(tmplFile)
	if err != nil {
		log.Fatalln(err)
	}

	//3create Dynamic routing

	//4Run server
	mux := defaultMux()

	router := dynamicRoutHandler(jsonMap, mux, tmp)

	port := ":8080"

	err = http.ListenAndServe(port, router)
	if err != nil {
		return
	}
	fmt.Printf("Open on port : %v", port)
}

func defaultMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "intro", http.StatusTemporaryRedirect)
	})
	return mux
}
func dynamicRoutHandler(paths map[string]Story, fallback http.Handler, tmp *template.Template) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		pathWithoutSlash, _ := strings.CutPrefix(request.URL.Path, "/")
		if value, exists := paths[pathWithoutSlash]; exists {
			err := tmp.Execute(writer, value)
			if err != nil {
				log.Print(err)
			}
		}
		fallback.ServeHTTP(writer, request)
	}
}
