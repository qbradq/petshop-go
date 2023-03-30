package petshopd

// Pet represents a single pet in the database.
type Pet struct {
	ID          int    // This pet's ID number
	Name        string // The pet's name
	Description string // A brief description of the pet
}
