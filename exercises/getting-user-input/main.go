package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/getting-user-input/note"
)

func main() {
	noteTitle, noteContent := getNodeData()

	note, err := note.New(noteTitle, noteContent)

	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()
	err = note.Save()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Saved successfully")
}

func getNodeData() (string, string) {
	noteTitle := getUserInput("Note title: ")
	noteContent := getUserInput("Note content: ")

	return noteTitle, noteContent
}

func getUserInput(text string) string {
	fmt.Print(text)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.Trim(text, "\n")
	text = strings.Trim(text, "\r")

	return text
}
