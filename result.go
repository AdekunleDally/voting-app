package main

import (
	"html/template"
	"net/http"
)

var resultTmpl = template.Must(template.ParseFiles("static/result.html"))

func ResultHandler(w http.ResponseWriter, r *http.Request) {
	var catCount, dogCount int
	dbConn.QueryRow(ctx, "SELECT count FROM votes WHERE option = 'cat'").Scan(&catCount)
	dbConn.QueryRow(ctx, "SELECT count FROM votes WHERE option = 'dog'").Scan(&dogCount)

	data := struct {
		Cats int
		Dogs int
	}{
		Cats: catCount,
		Dogs: dogCount,
	}

	err := resultTmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
