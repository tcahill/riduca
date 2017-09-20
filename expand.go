package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func expandURL(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	rows, err := db.Query("SELECT url FROM urls WHERE id = $1", id)
	if err != nil {
		render.Render(w, r, InvalidRequestError(err))
	}
	defer rows.Close()

	var url string
	for rows.Next() {
		err := rows.Scan(&url)
		if err != nil {
			render.Render(w, r, ServerError(err))
		}
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
