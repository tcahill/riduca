package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/render"
	"github.com/pkg/errors"
)

type ShortenURLRequest struct {
	LongURL string `json:"long_url"`
}

func (s *ShortenURLRequest) Bind(r *http.Request) error {
	_, err := url.Parse(s.LongURL)
	if err != nil {
		return errors.Wrap(err, "invalid URL")
	}
	return nil
}

type ShortenURLResponse struct {
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}

func (s *ShortenURLResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func shortenURL(w http.ResponseWriter, r *http.Request) {
	req := ShortenURLRequest{}
	render.Bind(r, &req)

	t := time.Now().String()
	h := sha1.Sum(append([]byte(req.LongURL), []byte(t)...))
	id := base64.URLEncoding.EncodeToString(h[:])[:6]
	fmt.Printf("id: %s", id)
	_, err := db.Exec("INSERT INTO urls (id, url) VALUES ($1, $2)", id, req.LongURL)
	if err != nil {
		render.Render(w, r, InvalidRequestError(err))
		return
	}

	shortURL := fmt.Sprintf("%s/%s", domainName, id)
	render.Render(w, r, &ShortenURLResponse{LongURL: req.LongURL, ShortURL: shortURL})
}
