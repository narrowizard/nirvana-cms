package models

type User struct {
	Model
	Account  string
	Password string
	Salt     string
}
