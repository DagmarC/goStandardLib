package reflection

import (
	"fmt"
	"goStandardLib/reflection/media"
)

func MovieApp() {
	fmt.Println("\n----REFLECTION, CUSTOM TYPES----")

	// Created obj the Old way - copies
	//myFavoriteMovie := media.NewMovie("HP 6", media.G, 89.9)

	// Created as media Catalogable interface
	var myFavoriteMovie media.Catalogable = &media.Movie{}
	myFavoriteMovie.NewMovie("HP 6", media.G, 89.9)

	fmt.Println(myFavoriteMovie.GetTitle(), myFavoriteMovie.GetRating())

	myFavoriteMovie.SetTitle("HP 1")
	fmt.Println(myFavoriteMovie.GetTitle(), myFavoriteMovie.GetRating())

}
