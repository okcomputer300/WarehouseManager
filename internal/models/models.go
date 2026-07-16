package models

import "database/sql"

type Location struct {
	ID   int
	Name string
}

type Tag struct {
	ID   int
	Name string
}

type Item struct {
	ID           int
	Name         string
	LocationID   int
	Amount       int
	Description  string
	LocationName string
	TagID        sql.NullInt64
	TagName      sql.NullString
}
