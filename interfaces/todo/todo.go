package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	filename := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("the todo text is required")
	}

	return Todo{text}, nil
}