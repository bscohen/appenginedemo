package javascript

import (
	"html/template"
	"net/http"

	"appengine"
	"appengine/user"
)

func init() {
	http.HandleFunc("/javascript", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if err := loginTemplate.Execute(w, u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var loginTemplate = template.Must(template.ParseFiles("templates/base.html",
	"javascript/templates/javascript.html"))
