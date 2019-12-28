package model

// TableDetails provides structural details about a table in postgres
type TableDetails struct {
	Table        string
	Id           string
	ArrayColumns map[string]struct{}
}
