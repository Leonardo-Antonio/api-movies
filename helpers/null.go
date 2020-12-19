package helpers

import "database/sql"

func NullString(value string) sql.NullString {
	null := sql.NullString{String: value}
	if null.String != "" {
		null.Valid = true
	}
	return null
}
