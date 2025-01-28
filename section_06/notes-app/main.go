package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes-app/note"
	"example.com/notes-app/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {

	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	_ = outputData(userNote)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving failed...")
		return err
	}

	fmt.Println("Saving succeded!")
	return nil
}

func getTodoData() string {
	return getUserInput("Todo Text: ")
}

func getNoteData() (string, string) {

	title := getUserInput("Note Title: ")
	content := getUserInput("Note Content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	var value, err = reader.ReadString('\n')

	if err != nil {
		return ""
	}

	return strings.TrimSuffix(value, "\n")
}
