package main

import (
	//	"fmt"

	//	"example.com/bank/fileops"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

type Player struct {
	name string
}

const balenceFileName = "balance.txt"


func main () {
	// var answer int;
	// balance, err := fileops.ReadBalanceFromFile();

	// if (err != nil) {
	// 	fmt.Println("ERROR")
	// 	fmt.Println(err)
	// 	fmt.Println("------------------------------")
	// }

	// fmt.Println("Welcome to Bankay the best banking account!");
	// fmt.Println("What do you want to do?");

	// for {
	// 	fmt.Println("1. Chack bank balance");
	// 	fmt.Println("2. Deposit money" );
	// 	fmt.Println("3. Whithdrawl money" );
	// 	fmt.Println("4. Exit");
	// 	fmt.Print("Enter Your Answer: ");
	// 	fmt.Scanln(&answer);

	// 	fmt.Println("###################################")

	// 	switch answer {
	// 		case 1:
	// 			fmt.Printf("Your Balance is %.2f\n", balance);
	// 		case 2:
	// 			var deposit float64;
	// 			fmt.Print("Your deposit: ")
	// 			fmt.Scan(&deposit);
	// 			balance += deposit;
	// 			fmt.Println("Your balance is: ", balance);
	// 			fileops.WriteBalanceToFile(balance)
	// 		case 3:
	// 			var whithdrawl float64;
	// 			fmt.Print("Whithdrawl amount: ")
	// 			fmt.Scan(&whithdrawl);
	// 			balance -= whithdrawl;
	// 			fmt.Println("Your balance is: ", balance);
	// 			fileops.WriteBalanceToFile(balance)
	// 		case 4:
	// 			fmt.Println("############# Goodbye #############");
	// 			fmt.Println("###################################")
	// 			return
	// 		default:
	// 			fmt.Println("The entred number is not valid");
	// 	}
	// }

    // fmt.Println(randomdata.FirstName(randomdata.Female))
    fmt.Println(randomdata.Email())
}

func isPlayerWining (player *Player) {
	player.name = "Player one"
}
