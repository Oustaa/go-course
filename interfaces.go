// // // // // // package main

// // // // // // import (
// // // // // // 	"os"
// // // // // // 	"time"
// // // // // // )

// // // // // // type Speaker interface {
// // // // // // 	speak(message string) string
// // // // // // }

// // // // // // type FileInfo interface {
// // // // // // 	Name() string
// // // // // // 	Size() int64
// // // // // // 	Mode() os.FileMode
// // // // // // 	ModTime() time.Time
// // // // // // 	IsDir() bool
// // // // // // 	Sys() interface{}
// // // // // // }

// // // // // // type Person34 struct {
// // // // // // }

// // // // // // func (p Person34) speak(message string) string {
// // // // // // 	return message
// // // // // // }

// // // // // // func initSpeaker(speaker Speaker) {}

// // // // // // func main() {
// // // // // // 	me := Person34{}
// // // // // // 	initSpeaker(me)
// // // // // // }

// // // // // type Speaker interface {
// // // // // 	Speak() string
// // // // // }
// // // // // type cat struct {
// // // // // }

// // // // // func main() {
// // // // // 	c := cat{}
// // // // // 	fmt.Println(c.Speak())
// // // // // 	c.Greeting()
// // // // // }
// // // // // func (c cat) Speak() string {
// // // // // 	return "Purr Meow"
// // // // // }
// // // // // func (c cat) Greeting() {
// // // // // 	fmt.Println("Meow,Meow!!!!mmmeeeeoooowwww")
// // // // // }

// // // // // package main

// // // // // import (
// // // // // 	"fmt"
// // // // // )

// // // // // type Speaker interface {
// // // // // 	Speak() string
// // // // // }
// // // // // type cat struct {
// // // // // 	name string
// // // // // 	age  int
// // // // // }

// // // // // func main() {
// // // // // 	c := cat{name: "Oreo", age: 9}
// // // // // 	fmt.Println(c.Speak())
// // // // // 	fmt.Println(c)
// // // // // }
// // // // // func (c cat) Speak() string {
// // // // // 	return "Purr Meow"
// // // // // }

// // // // // func (c cat) String() string {
// // // // // 	return fmt.Sprintf("%v (%v years old)", c.name, c.age)
// // // // // }

// // // // package main

// // // // import "fmt"

// // // // type Speaker interface {
// // // // 	Speak() string
// // // // }

// // // // type person struct {
// // // // 	name      string
// // // // 	age       int
// // // // 	isMarried bool
// // // // }

// // // // func main() {
// // // // 	p := person{name: "Cailyn", age: 44, isMarried: false}
// // // // 	fmt.Println(p.Speak())
// // // // 	fmt.Println(p)
// // // // }

// // // // func (p person) String() string {
// // // // 	return fmt.Sprintf("%v (%v years old).\nMarried status: %v ",
// // // // 		p.name, p.age, p.isMarried)
// // // // }

// // // // func (p person) Speak() string {
// // // // 	return "Hi my name is: " + p.name
// // // // }

// // // // package main

// // // // import (
// // // // 	"fmt"
// // // // )

// // // // type Speaker interface {
// // // // 	Speak() string
// // // // }

// // // // func saySomething(say ...Speaker) {
// // // // 	for _, s := range say {
// // // // 		fmt.Println(s.Speak())
// // // // 	}
// // // // }

// // // // type cat struct{}

// // // // func (c cat) Speak() string {
// // // // 	return "Purr Meow"
// // // // }

// // // // type dog struct{}

// // // // func (d dog) Speak() string {
// // // // 	return "Woof Woof"
// // // // }

// // // // type person struct {
// // // // 	name string
// // // // }

// // // // func (p person) Speak() string {
// // // // 	return "Hi my name is " + p.name + "."
// // // // }
// // // // func main() {
// // // // 	c := cat{}
// // // // 	d := dog{}
// // // // 	p := person{name: "Heather"}
// // // // 	saySomething(c, d, p)
// // // // }

// // // package main

// // // import "fmt"

// // // type Shape interface {
// // // 	area() float64
// // // 	name() string
// // // }

// // // type triangle struct {
// // // 	base   float64
// // // 	height float64
// // // }
// // // type rectangle struct {
// // // 	length float64
// // // 	width  float64
// // // }
// // // type square struct {
// // // 	side float64
// // // }

// // // func (t *triangle) area() float64 {
// // // 	return t.base * t.height / 2
// // // }

// // // func (t *triangle) name() string {
// // // 	return "triangle"
// // // }

// // // func (r rectangle) Area() float64 {
// // // 	return r.length * r.width
// // // }

// // // func (r rectangle) name() string {
// // // 	return "rectangle"
// // // }

// // // func (s square) Area() float64 {
// // // 	return s.side * s.side
// // // }
// // // func (s square) Name() string {
// // // 	return "square"
// // // }

// // //	func printShapeDetails(shapes ...Shape) {
// // //		for _, shape := range shapes {
// // //			fmt.Printf("%s has an area of %.2f", shape.name(), shape.area())
// // //		}
// // //	}
// // package main

// // import (
// // 	"encoding/json"
// // 	"fmt"
// // 	"io"
// // 	"strings"
// // )

// // type Person struct {
// // 	Name string `json:"name"`
// // 	Age  int    `json:"age"`
// // }

// // func main() {
// // 	s := `{"Name":"Joe","Age":18}`
// // 	s2 := `{"Name":"Jane","Age":21}`
// // 	p, err := loadPerson(strings.NewReader(s))
// // 	if err != nil {
// // 		fmt.Println(err)
// // 	}
// // 	fmt.Println(p)
// // 	p2, err := loadPerson2(s2)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 	}
// // 	fmt.Println(p2)
// // }

// // func loadPerson2(s string) (Person, error) {
// // 	var p Person
// // 	err := json.NewDecoder(strings.NewReader(s)).Decode(&p)
// // 	if err != nil {
// // 		return p, err
// // 	}
// // 	return p, nil
// // }

// // func loadPerson(r io.Reader) (Person, error) {
// // 	var p Person
// // 	err := json.NewDecoder(r).Decode(&p)
// // 	if err != nil {
// // 		return p, err
// // 	}
// // 	return p, err
// // }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Printf("an error occured, %v\n", err)
// 		}
// 	}()

// 	var str interface{} = 2
// 	v := str.(string)
// 	fmt.Println(strings.ToTitle(v))
// }

package main

import (
	"fmt"
)

func main() {
	var str interface{} = "the book club"
	v, isValid := str.(int)
	if isValid {
		fmt.Println(v)
	} else {
		fmt.Println("You must provide a number")
	}
}

func multiTypeFunciton(value interface{}) {

	switch value := value.(type) {
	case string:
		fmt.Printf("%s is of type string\n", value)
	}

}
