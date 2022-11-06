package database

import (
	"errors"
	"time"
)

// User -
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

// Create User
func (c Client) CreateUser(
	email, password, name string,
	age int,
) (User, error) {
	dat, err := c.readDB()
	if err != nil {
		return User{}, err
			}
	// Check if email is in the users list
	_, ok := dat.Users[email]
	if ok {
		return User{}, errors.New("User already exists! Cannot create.")
	}
	createdUser := User{
		CreatedAt: time.Now().UTC(),
		Email: email,
		Password: password,
		Name: name,
		Age: age,
	}
	// Add to database
	dat.Users[email] = createdUser
	// Write to database
	err = c.updateDB(dat)
	if err != nil {
		return User{}, err
	}
	// Create a new user
	return createdUser, err
}

func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
	dat, err := c.readDB()
	if err != nil {
		return User{}, err
			}
	// Check if email is in the users list
	existinguser, ok := dat.Users[email]
	if !ok {
		return User{}, errors.New("user doesn't exist")
	}
	existinguser.Password = password
	existinguser.Name = name
	existinguser.Age = age
	dat.Users[email] = existinguser
	// Write to database
	err = c.updateDB(dat)
	if err != nil {
		return User{}, err
	}
	// Create a new user
	return existinguser, err
}

func (c Client) GetUser(email string) (User, error) {
	dat, err := c.readDB()
	if err != nil {
		return User{}, err
			}
	existinguser, ok := dat.Users[email]
	if !ok {
		return User{}, errors.New("user does not exist")
	}
	return existinguser, nil
}

func (c Client) DeleteUser(email string) error {
	dat, err := c.readDB()
	if err != nil {
		return err
			}
	_, ok := dat.Users[email]
	if !ok {
		return errors.New("nothing to delete, user does not exist")
	}
	delete(dat.Users, email)
	err = c.updateDB(dat)
	if err != nil {
		return err
	}
	return nil
}
