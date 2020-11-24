package main

import (
	"encoding/json"
	"fmt"
)

type (
	Note struct {
		Title string
		Body  string
	}

	Notes []Note
)

func getNotes() Notes {
	notes := make(Notes, 0)
	json.Unmarshal(readFile(FileName), &notes)
	return notes
}

func createNote(note Note) {

	collection := append(getNotes(), Note{
		Title: note.Title,
		Body:  note.Body,
	})

	content, error := json.Marshal(collection)

	throwErrorIfNecessary(error)

	writeToFile(content)
}

func renderNotes(loadedNotes Notes) {
	for _, note := range loadedNotes {
		fmt.Println("====================")
		fmt.Println("Title: " + note.Title)
		fmt.Println("Body: " + note.Body)
	}
}