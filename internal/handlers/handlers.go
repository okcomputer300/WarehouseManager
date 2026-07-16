package handlers

import (
	"database/sql"
	"html/template"
	"lagermanager/internal/database"
	"lagermanager/internal/models"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ResolveTemplatePath() string {
	candidates := []string{
		filepath.Join("web", "templates", "index.html"),
		filepath.Join("templates", "index.html"),
		"/app/web/templates/index.html",
		"/app/templates/index.html",
	}

	for _, candidate := range candidates {
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}

	return candidates[0]
}

func ListItemsHandler(w http.ResponseWriter, r *http.Request) {
	query := `
        SELECT 
            i.id, i.name, i.location_id, i.amount, i.description, 
            l.name AS location_name, 
            i.tag_id, t.name AS tag_name
        FROM items i 
        LEFT JOIN locations l ON i.location_id = l.id
        LEFT JOIN tags t ON i.tag_id = t.id
    `
	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []models.Item

	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.LocationID, &item.Amount, &item.Description, &item.LocationName, &item.TagID, &item.TagName)
		if err != nil {
			log.Println("error scanning:", err)
			continue
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles(ResolveTemplatePath())
	if err != nil {
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, items)
	if err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
