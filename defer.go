package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota //starts at zero
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Dev struct {
	fullname string
	workWeek [7]int
}

func (d *Dev) addWOrkHours(day Weekday, hours int) {
	d.workWeek[day] = hours
}

func (d *Dev) HoursWorked() int {
	total := 0

	for _, hours := range d.workWeek {
		total += hours
	}

	return total
}

func _defered() {
	fmt.Println("This is the call of the deffered function")
}

func _fff() {
	defer _defered()
	fmt.Println("Oussama Syas Hello")
	return
}

func maisd234n() {
	_fff()

	me := Dev{"Oussama Tailba", [7]int{}}

	me.addWOrkHours(Monday, 7)
	me.addWOrkHours(Friday, 8)
	me.addWOrkHours(Saturday, 4)
	me.addWOrkHours(Sunday, 0)
	me.addWOrkHours(Thursday, 8)
	me.addWOrkHours(Tuesday, 8)
	me.addWOrkHours(Wednesday, 8)

	fmt.Printf("Total work hours is %v\n", me.HoursWorked())

}
