package models

import "database/sql"

type Location struct {
	ID          int
	Name        string
	Description string
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
	PhotoPath    sql.NullString
}

func (i Item) HasTag() bool {
	return i.TagID.Valid && i.TagID.Int64 > 0
}

func (i Item) HasPhoto() bool {
	return i.PhotoPath.Valid && i.PhotoPath.String != ""
}
