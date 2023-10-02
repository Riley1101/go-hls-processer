package src

type Movie struct {
	ID          int
	Title       string
	Description string
	File        File
}

func NewMovie(title string, description string, file File) *Movie {
	return &Movie{
		Title:       title,
		Description: description,
		File:        file,
	}
}
