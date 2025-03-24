package main

import "fmt"

type Computer struct {
	CPU   string
	GPU   string
	Games []string
}

func (c *Computer) addGames(game ...string) {
	c.Games = append(c.Games, game...)
}

func main() {
	myPc := Computer{"I7 12000K", "RTX 4080", []string{}}

	fmt.Printf("%+v\n", myPc)

	myPc.addGames("CS2", "ARK Survival")

	fmt.Printf("%#v\n", myPc)

}
