package internal

import "fmt"

type FilmID string
type FilmName string

// FilmDB is a map of FilmID to FilmName
var FilmDB map[FilmID]FilmName

// NewFilmDB creates a new film database
func NewFilmDB() map[FilmID]FilmName {
	if FilmDB == nil {
		FilmDB = make(map[FilmID]FilmName)
	}
	return FilmDB
}

// AddFilm adds a new film to the database
func AddFilm(id, name string) error {
	if id == "" || name == "" {
		return fmt.Errorf("invalid film data: FilmID, FilmName cannot be empty")
	}
	NewFilmDB()
	if _, ok := FilmDB[FilmID(id)]; !ok {
		FilmDB[FilmID(id)] = FilmName(name)
	} else {
		return fmt.Errorf("film already exists")
	}
	return nil
}

// RemoveFilm removes a film from the database
func RemoveFilm(id string) error {
	if id == "" {
		return fmt.Errorf("invalid film data: FilmID cannot be empty")
	}
	if _, ok := FilmDB[FilmID(id)]; ok {
		delete(FilmDB, FilmID(id))
		RemoveAuthorizationForFilm(id)
	} else {
		return fmt.Errorf("film not found")
	}
	return nil
}

// IsValidFilm checks if the FilmID is valid
func IsValidFilm(id string) bool {
	isValid := false
	if id == "" || FilmDB == nil {
		return false
	}
	if _, ok := FilmDB[FilmID(id)]; ok {
		isValid = true
	}
	return isValid
}
