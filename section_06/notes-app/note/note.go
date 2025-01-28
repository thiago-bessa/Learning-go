package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Println(note.Title, "-", note.CreatedAt)
	fmt.Println(note.Content)
}

func (note Note) Save() error {
	var filename = strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	var json, err = json.Marshal(note)

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, json, 0644)

	if err != nil {
		return err
	}

	return nil
}
