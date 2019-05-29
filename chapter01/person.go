package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Saver persists the supplied bytes
type Saver interface {
	Save(data []byte) error
}

// SavePerson will validate and persist the supplied person
func SavePerson(person *Person, saver Saver) error {
	// validate the inputs
	err := person.validate()
	if err != nil {
		return err
	}

	// encode person to bytes
	bytes, err := person.encode()
	if err != nil {
		return err
	}

	// save the person and return the result
	return saver.Save(bytes)
}

// Person data object
type Person struct {
	Name  string
	Phone string
}

// validate the person object
func (p *Person) validate() error {
	if p.Name == "" {
		return errors.New("name missing")
	}

	if p.Phone == "" {
		return errors.New("phone missing")
	}
	return nil
}

// convert person into bytes
func (p *Person) encode() ([]byte, error) {
	return json.Marshal(p)
}

// load person by ID
func loadPerson(ID int) ([]byte, error) {
	return nil, nil
}

// LoadPerson will load the requested person by ID.
// Errors include: invalid ID, missing person and failure to load
// or decore
func LoadPerson(ID int, decodePerson func(data []byte) *Person) (*Person, error) {
	// validate the input
	if ID <= 0 {
		return nil, fmt.Errorf("invalid ID '%d' supplied", ID)
	}

	// load from storage
	bytes, err := loadPerson(ID)
	if err != nil {
		return nil, err
	}

	// decode bytes and return
	return decodePerson(bytes), nil
}
