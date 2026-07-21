package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"lagermanager/internal/database"
	"lagermanager/internal/handlers"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func resolveExistingDir(candidates ...string) string {
	for _, candidate := range candidates {
		if info, err := os.Stat(candidate); err == nil && info.IsDir() {
			return candidate
		}
	}
	return candidates[0]
}

func resolveTemplatesDir() string {
	cwd, _ := os.Getwd()
	return resolveExistingDir(
		filepath.Join(cwd, "web", "templates"),
		filepath.Join(cwd, "templates"),
		"web/templates",
		"templates",
		"/app/web/templates",
		"/app/templates",
	)
}

func resolveStaticDir() string {
	cwd, _ := os.Getwd()
	return resolveExistingDir(
		filepath.Join(cwd, "web", "static"),
		filepath.Join(cwd, "static"),
		"web/static",
		"static",
		"/app/web/static",
		"/app/static",
	)
}

func main() {
	staticDir := resolveStaticDir()

	if err := database.InitDB(); err != nil {
		log.Fatalf("Error initialzing DB: %v", err)
	}

	defer database.DB.Close()

	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.ListItemsHandler)
	http.HandleFunc("/modify-item", handlers.ModifyItemHandler)
	http.HandleFunc("/add-item", handlers.AddItemHandler)

	http.Handle("/metrics", promhttp.Handler())

	log.Println("server works on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
