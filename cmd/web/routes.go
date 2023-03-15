package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	
	return standard.Then(mux)
}

// You don’t need to use justinas/alice package, but the reason I recommend it is because it makes it easy
// to create composable, reusable, middleware chains — and that can be a real help as your
// application grows and your routes become more complex. The package itself is also small and
// lightweight, and the code is clear and well written.