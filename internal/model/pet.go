package model

import "image"

// Pet represents a pet for adoption.
type Pet struct {
	Name        string      // Name is the pet's name
	Description string      // Description is a short description of the pet
	Picture     image.Image // Picture should be a photo of the pet
}

//
