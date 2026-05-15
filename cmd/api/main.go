package main

import (
	"encoding/json"
	"net/http"

	"github.com/2spy/vinted-discord-bot/internal/scrapers/vinted"
	"github.com/2spy/vinted-discord-bot/pkg/models"
)

func main() {

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")

		query := r.URL.Query().Get("q")
                size := r.URL.Query().Get("size")
                category := r.URL.Query().Get("category")

		scraper := vinted.NewVintedScraper()

		items, err := scraper.Search(r.Context(), models.ScrapeJob{
			Query: query,
                        Size: size,
                        Category: category,
		})

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	})

	http.ListenAndServe(":8080", nil)
}