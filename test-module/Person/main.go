package person

type Location struct {
	street  string
	city    string
	country string
}

type Person struct {
	Name string
	Age  int
	Location
}

func New(name string, age int, street, city, country string) *Person {
	myPerson := Person{name, age, Location{street, city, country}}

	return &myPerson
}
