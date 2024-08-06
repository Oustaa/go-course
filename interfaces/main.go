package main

import "fmt"

// type saver interface {
// 	Save() error
// }

// type outputtable interface {
// 	saver
// 	Display()
// }

func main() {

	result := add("1", "2")

	fmt.Println(result + "1");

	// fmt.Println("Welcome to note-traker")

	// title := getData("Enter note title: ")
	// content := getData("Enter note content: ")
	// todoText := getData("Enter todo text: ")

	// myTodo, _ := todo.New(todoText)
	// n, _ := note.New(title, content)

	// savedata(myTodo)

	// savedata(n)
}

// func printSomthiig(value interface{}) {
// 	switch value.(type) {
// 		case string:
// 			fmt.Println("THIS IS A STRING")
// 		case float32:
// 		case float64:
// 			fmt.Println("THIS IS A FLOAT")
// 	}
// }

// func savedata(t saver) {
// 	err := t.Save()

// 	if err != nil {
// 		fmt.Println("Error while saving data")
// 	}

// 	fmt.Println("data saved successfully")
// }

// func getData(msg string) string {
// 	fmt.Print(msg)

// 	redear := bufio.NewReader(os.Stdin)
// 	input, _ := redear.ReadString('\n')

// 	input = strings.TrimSuffix(input, "\n")
// 	input = strings.TrimSuffix(input, "\r")

// 	return input
// }

func add[T string | float32 | float64 | int](a, b T) T {
	return a + b
	// aInt, aIsInt := a.(int)
	// bInt, bIsInt := b.(int)

	// if aIsInt && bIsInt {
	// 	return aInt + bInt
	// }

	// aFloat, aIsFloat := a.(float64)
	// bFloat, bIsFloat := b.(float64)

	// if aIsFloat && bIsFloat {
	// 	return aFloat + bFloat
	// }

	// aString, aIsString := a.(string)
	// bString, bIsString := b.(string)

	// if aIsString && bIsString {
	// 	return aString + bString
	// }

	// return nil
}