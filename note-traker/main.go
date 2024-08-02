package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/noteTracker/note"
)

func main() {
	fmt.Println("Welcome to note-traker")

	title := getData("Enter note title: ")
	content := getData("Enter note content: ")

	n, error := note.New(title, content)

	if error != nil {
		fmt.Println(error)
		return
	}

	n.Display()
	err := n.Save()

	if err != nil {
		fmt.Println(err)
	}
}

func getData(msg string) string {
	fmt.Print(msg)

	redear := bufio.NewReader(os.Stdin)
	input, _ := redear.ReadString('\n')

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}