package media

import "strings"

type Catalogable interface {
	NewMovie(title string, rating Rating, boxOffice float32) // constructor
	GetTitle() string
	GetRating() Rating
	GetBoxOffice() float32
	SetTitle(newTitle string)
	SetRating(newRating Rating)
	SetBoxOffice(newBoxOffice float32)
}

type Movie struct {
	title     string
	rating    Rating
	boxOffice float32
}

type Rating string

const (
	R  = "R (Restricted)"
	G  = "G (General)"
	PG = "PG (Parental Guidance)"
)

// Returning a new copy of a movie
//func NewMovie(title string, rating Rating, boxOffice float32) Movie {
//	return Movie{
//		title:     title,
//		rating:    rating,
//		boxOffice: boxOffice,
//	}
//}

// NewMovie CONSTRUCTOR - keeping our movie as a struct, that can be modified directly
func (m *Movie) NewMovie(title string, rating Rating, boxOffice float32) {
	m.title = title
	m.rating = rating
	m.boxOffice = boxOffice
}

func (m Movie) GetTitle() string {
	return strings.ToTitle(m.title)
}

func (m *Movie) GetRating() Rating {
	return m.rating
}

func (m Movie) GetBoxOffice() float32 {
	return m.boxOffice
}

// SetTitle SETTERS MUST be POINTER BASED receivers, otherwise no change is reflected
func (m *Movie) SetTitle(title string) {
	m.title = title
}

func (m *Movie) SetRating(rating Rating) {
	m.rating = rating
}

func (m *Movie) SetBoxOffice(boxOffice float32) {
	m.boxOffice = boxOffice
}
