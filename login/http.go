package login

import (
	"html/template"
	"net/http"

	"appengine"
	"appengine/user"
)

func init() {
	http.HandleFunc("/login", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	if err := loginTemplate.Execute(w, u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var loginTemplate = template.Must(template.ParseFiles("templates/base.html",
	"login/templates/login.html"))
