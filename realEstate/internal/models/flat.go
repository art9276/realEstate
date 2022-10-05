package models

type Flat struct {
	ID    int
	Price int
}

// GetID returns the user ID.
func (f Flat) GetID() string {

	return string("Напиши меня")
}

// GetName returns the user name.
func (f Flat) GetPrice() string {
	return string("Напиши меня")
}
