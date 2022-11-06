package database

import (
	"encoding/json"
	"errors"
	"os"
)

type Client struct {
	dbPath string
}

type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

func NewClient(filename string) Client {
	return Client{
		dbPath: filename,
	}
}

func (c Client) EnsureDB() error {
	_, err := os.ReadFile(c.dbPath)
	if errors.Is(err, os.ErrNotExist) {
		return c.createDB()
	}
	return err
}


func (c Client) createDB() error {
	dat, err := json.Marshal(databaseSchema{
		Users: make(map[string]User),
		Posts: make(map[string]Post),
	})
	if err != nil {
	return err
	}
	err = os.WriteFile(c.dbPath, dat, 0600)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) updateDB(db databaseSchema) error {
	dat, err := json.Marshal(db)
	if err != nil {
	return err
	}
	err = os.WriteFile(c.dbPath, dat, 0600)
	if err != nil {
	return err
	}
	return nil
}

func (c Client) readDB() (databaseSchema, error) {
	readDBSchema := databaseSchema{}
	dat, err := os.ReadFile(c.dbPath)
	if errors.Is(err, os.ErrNotExist) {
		return readDBSchema, nil
	}
	nerr := json.Unmarshal(dat, &readDBSchema)
	if nerr != nil {
		return databaseSchema{}, nerr
	}
	return readDBSchema, nil
}
