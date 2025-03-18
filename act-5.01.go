package main

type employee struct {
	is        int
	firstName string
	lastName  string
}

type developer struct {
	individual employee
	HourlyRate int
	workWeek   [7]int
}

const (
	MONDAY = iota
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)
