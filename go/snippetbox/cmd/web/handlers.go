package main

import (
	"fmt"
	"net/http"
	"snippetbox/internal/validator"
	"strconv"
)

type snippetCreateForm struct {
	Title   string
	Content string
	Expires int
	validator.Validator
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.tmpl.html", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.serverError(w, r, err)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		app.clientError(w, http.StatusNotFound)
		return
	}

	flash := app.sessionManager.PopString(r.Context(), "flash")

	data := app.newTemplateData(r)
	data.Snippet = s
	data.Flash = flash

	app.render(w, r, http.StatusOK, "view.tmpl.html", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	form := snippetCreateForm{
		Expires: 7,
	}

	data.Form = form

	app.render(w, r, http.StatusOK, "create.tmpl.html", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	var form snippetCreateForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.Validator.CheckField(validator.NotBlank(form.Title), "title", "field must not be blank")
	form.Validator.CheckField(validator.MaxChars(form.Title, 100), "title", "field must not be larger tnan 100 characters")

	form.Validator.CheckField(validator.NotBlank(form.Content), "content", "field must not be blank")
	form.Validator.CheckField(validator.MaxChars(form.Content, 1000), "content", "field must not be larger tnan 100 characters")

	form.Validator.CheckField(validator.PermittedValue(form.Expires, 1, 7, 365), "expires", "value must be either 1, 7 or 365")

	if !form.Validator.IsValid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "create.tmpl.html", data)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
