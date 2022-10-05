package models

type House struct {
	ID    int
	Price int
}

// GetID returns the user ID.
func (h House) GetID() string {

	return string("Напиши меня")
}

// GetName returns the user name.
func (h House) GetPrice() string {
	return string("Напиши меня")
}
