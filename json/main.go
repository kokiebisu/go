package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

// Person is a unit of a Person
type Person struct {``
	First string `json:"First"`
	Last string `json:"Last"`
}

type Data struct {
	People []Person
}

func main() {
	http.HandleFunc("/to", ToJSON)
	http.HandleFunc("/from", FromJSON)
	http.ListenAndServe(":8080", nil)
}

// ToJSON converts struct into JSON and writes to the response writer
func ToJSON(w http.ResponseWriter, req *http.Request) {
	p := Data{
		People: []Person{
			{"ken", "okiebisu"}, 
			{"john", "park"},
		},
	}
	j, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintln(w, string(j))
	// io writestring only accepts type string
	// io.WriteString()
}

// FromJSON converts the JSON back to the person struct
func FromJSON(w http.ResponseWriter, req *http.Request) {
	var d Data
	j := []byte(`{"People": [{"First": "John", "Last": "Park"}, {"First": "Ken", "Last": "Okiebisu"}]}`)
	err := json.Unmarshal(j, &d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(d.People)
	// fmt.Fprintln(w, p)
	tpl.ExecuteTemplate(w, "index.html", d)
}