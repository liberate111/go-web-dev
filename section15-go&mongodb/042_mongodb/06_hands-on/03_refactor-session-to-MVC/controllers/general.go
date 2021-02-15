package controllers

import (
	"html/template"
	"mvc/session"
	"net/http"
)

type Controller struct {
	tpl *template.Template
}

func NewController(t *template.Template) *Controller {
	return &Controller{t}
}

func (c Controller) Index(w http.ResponseWriter, req *http.Request) {
	u := session.GetUser(w, req)
	session.Show()
	c.tpl.ExecuteTemplate(w, "index.gohtml", u)

}

func (c Controller) Bar(w http.ResponseWriter, req *http.Request) {
	u := session.GetUser(w, req)
	if !session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter this site", http.StatusForbidden)
		return
	}
	session.Show()
	c.tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
