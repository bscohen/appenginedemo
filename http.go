package guestbook

import (
	"html/template"
	"net/http"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

type User struct {
	Email    string
	LoggedIn bool
}

func init() {
	http.HandleFunc("/", root)
}

// guestbookKey returns the key used for all guestbook entries.
func guestbookKey(c appengine.Context) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	ud := &User{LoggedIn: false}
	if u != nil {
		ud.Email = u.Email
		ud.LoggedIn = true
	}
	if err := homeTemplate.Execute(w, ud); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var homeTemplate = template.Must(template.ParseFiles("templates/base.html",
	"templates/home-content.html"))
