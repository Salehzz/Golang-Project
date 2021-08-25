package main

import (
	"net/http"
)

//all informations about user
type ContactDetails struct {
	Name    string
	Email   string
	Subject string
	Message string
}

//form handler for renders the form.html or formshow.html
func form(w http.ResponseWriter, r *http.Request) {
	pageVars := PageVars{
		Title: "Forms",
	}
	if r.Method != http.MethodPost {
		render(w, "form.html", pageVars)
		return
	}
	details := ContactDetails{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}

	// use informations , we got them from user
	pageVars = PageVars{
		Title: ":)",
		Name:  details.Name,
	}
	render(w, "formshow.html", pageVars)
}
