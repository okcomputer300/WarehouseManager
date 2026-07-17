package handlers

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"lagermanager/internal/database"
	"lagermanager/internal/models"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

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
	searchTerm := r.URL.Query().Get("search")
	query := `
        SELECT
            i.id, i.name, i.location_id, i.amount, i.description,
            l.name AS location_name,
            i.tag_id, t.name AS tag_name
        FROM items i
        LEFT JOIN locations l ON i.location_id = l.id
        LEFT JOIN tags t ON i.tag_id = t.id
    `

	var rows *sql.Rows
	var err error

	if searchTerm != "" {
		query += ` WHERE i.name LIKE ?
					  OR i.description LIKE ?
					  OR l.name LIKE ?
					  OR t.name LIKE ?`

		searchParam := "%" + searchTerm + "%"
		rows, err = database.DB.Query(query, searchParam, searchParam, searchParam, searchParam)
	} else {
		rows, err = database.DB.Query(query)
	}

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
			log.Println("error scanning item:", err)
			continue
		}
		items = append(items, item)
	}

	locationsQuery := `SELECT id, name, description FROM locations ORDER BY name`
	locationRows, err := database.DB.Query(locationsQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer locationRows.Close()

	var locations []models.Location
	for locationRows.Next() {
		var location models.Location
		err := locationRows.Scan(&location.ID, &location.Name, &location.Description)
		if err != nil {
			log.Println("error scanning location:", err)
			continue
		}
		locations = append(locations, location)
	}

	tagsQuery := `SELECT id, name FROM tags ORDER BY name`
	tagRows, err := database.DB.Query(tagsQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tagRows.Close()

	var tags []models.Tag
	var tagNames []string
	for tagRows.Next() {
		var tag models.Tag
		err := tagRows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			log.Println("error scanning tag:", err)
			continue
		}
		tags = append(tags, tag)
		tagNames = append(tagNames, tag.Name)
	}

	tagsJSON, err := json.Marshal(tagNames)
	if err != nil {
		http.Error(w, "Failed to encode tags", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles(ResolveTemplatePath())
	if err != nil {
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Items     []models.Item
		Locations []models.Location
		Tags      []models.Tag
		TagsJSON  template.JS
		Search    string
	}{
		Items:     items,
		Locations: locations,
		Tags:      tags,
		TagsJSON:  template.JS(tagsJSON),
		Search:    searchTerm,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func ModifyItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	itemID := r.FormValue("id")
	action := r.FormValue("action")
	amountStr := r.FormValue("amount")
	tagName := r.FormValue("tag")

	if itemID == "" {
		http.Error(w, "Missing item id", http.StatusBadRequest)
		return
	}

	if action == "tag" {
		if tagName == "" {
			http.Error(w, "Missing tag", http.StatusBadRequest)
			return
		}

		var tagID int
		err := database.DB.QueryRow(`SELECT id FROM tags WHERE name = ?`, tagName).Scan(&tagID)
		if err == sql.ErrNoRows {
			res, err := database.DB.Exec(`INSERT INTO tags (name) VALUES (?)`, tagName)
			if err != nil {
				log.Println("Error creating tag:", err)
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}
			tagID64, err := res.LastInsertId()
			if err != nil {
				log.Println("Error reading tag id:", err)
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}
			tagID = int(tagID64)
		} else if err != nil {
			log.Println("Error querying tag:", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		_, err = database.DB.Exec(`UPDATE items SET tag_id = ? WHERE id = ?`, tagID, itemID)
		if err != nil {
			log.Println("Error assigning tag:", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"ok": true, "tag": tagName})
		return
	}

	if action == "clear-tag" {
		_, err := database.DB.Exec(`UPDATE items SET tag_id = NULL WHERE id = ?`, itemID)
		if err != nil {
			log.Println("Error clearing tag:", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"ok": true, "tag": ""})
		return
	}

	if action == "" || amountStr == "" {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	step, err := strconv.Atoi(amountStr)
	if err != nil || step <= 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	var updateQuery string
	if action == "increase" {
		updateQuery = `UPDATE items SET amount = GREATEST(0, amount + ?) WHERE id = ?`
	} else if action == "decrease" {
		updateQuery = `UPDATE items SET amount = GREATEST(0, amount - ?) WHERE id = ?`
	} else {
		http.Error(w, "Invalid action", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(updateQuery, step, itemID)
	if err != nil {
		log.Println("Error updating item:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	var newAmount int
	err = database.DB.QueryRow("SELECT amount FROM items WHERE id = ?", itemID).Scan(&newAmount)
	if err != nil {
		log.Println("Error reading updated amount:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"ok":     true,
		"amount": newAmount,
	})
}
