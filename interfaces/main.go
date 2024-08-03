package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/interfaces/note"
	"example.com/interfaces/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	fmt.Println("Welcome to note-traker")

	title := getData("Enter note title: ")
	content := getData("Enter note content: ")
	todoText := getData("Enter todo text: ")

	myTodo, _ := todo.New(todoText)
	n, _ := note.New(title, content)

	savedata(myTodo)


	savedata(n)
}

func outputData() {

}

func savedata(t saver) {
	err := t.Save()

	if err != nil {
		fmt.Println("Error while saving data")
	}

	fmt.Println("data saved successfully")
}

func getData(msg string) string {
	fmt.Print(msg)

	redear := bufio.NewReader(os.Stdin)
	input, _ := redear.ReadString('\n')

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}