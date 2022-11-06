package database

import (
	"fmt"
	"time"
	"errors"
	"github.com/google/uuid"
)

// Post -
type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

func (c Client) CreatePost(userEmail, text string) (Post, error) {
	readDatabase, err := c.readDB()
	if err != nil {
		fmt.Println("Error reading database")
		return Post{}, err
	}
	_, err = c.GetUser(userEmail)
	if err != nil {
		fmt.Println("Error reading the user")
		return Post{}, err
	}
	id := uuid.New().String()
	// New post structure
	newPost := Post{
		ID: id,
		CreatedAt: time.Now().UTC(),
		UserEmail: userEmail,
		Text: text,
	}
	// Add post and update db
	readDatabase.Posts[id] = newPost
	c.updateDB(readDatabase)
	return newPost, nil
}

func (c Client) GetPosts(userEmail string) ([]Post, error) {
	readDatabase, err := c.readDB()
	if err != nil {
		return []Post{}, err
	}
	postsSlice := []Post{}
	for _, eachPost := range(readDatabase.Posts) {
		if eachPost.UserEmail == userEmail {
			postsSlice = append(postsSlice, eachPost)
			}
	}
	return postsSlice, nil
}

func (c Client) DeletePost(id string) error {
	// Read the db and iterate through each user
	db, err := c.readDB()
	if err != nil {
		return err
	}
	if _, ok := db.Posts[id]; !ok {
		return errors.New("Nothing to delete, post does not exist")
	}
	delete(db.Posts, id)
	err = c.updateDB(db)
	if err != nil {
		return err
	}
	return nil
}
