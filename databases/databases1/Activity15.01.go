package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	users := []struct {
		Name  string
		Email string
	}{
		{"Oussama Tailba", "Otailaba98@gmail.com"},
		{"Younes Tailba", "YounesTai6@gmail.com"},
	}

	db, err := sql.Open("postgres", "user=ousta password=Password12 host=localhost port=5433 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection to the database was successfully...")
	}

	createUsersTable := `
		CREATE TABLE public.users (
			id SERIAL PRIMARY KEY,
			name VARCHAR UNIQUE,
			email VARCHAR UNIQUE
		)
	`

	_, err = db.Exec(createUsersTable)
	if err != nil {
		fmt.Println("There was an error creating the table 'users'...")
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println("The table was created successfully...")
	}

	insertUserStm, err := db.Prepare(`
		INSERT INTO users (name, email) 
		values($1, $2)
	`)
	if err != nil {
		fmt.Println("There was an error preparing the insert steatment")
	}

	for _, user := range users {
		_, err = insertUserStm.Exec(user.Name, user.Email)
		if err != nil {
			fmt.Printf("There was an error inserting the user with email: %s\n", user.Email)
			fmt.Printf("%v\n", err)
		}
	}

	updateStmt, err := db.Prepare(`
		update users
		SET email = $1
		WHERE ID = $2
	`)
	if err != nil {
		fmt.Println("There is an error while preparing the update steatment")
		fmt.Printf("%v\n", err)
	}
	updateResult, updateResultErr := updateStmt.Exec("user@packt.com", 1)
	if updateResultErr != nil {
		fmt.Println("An error occured while updating the row")
		fmt.Printf("%v\n", updateResultErr)
	}

	updateResultCount, updateResultCountErr := updateResult.RowsAffected()
	if updateResultCountErr != nil {
		fmt.Println("Error Error Error Error Error")
		fmt.Printf("%v\n", updateResultCountErr)
	} else {
		fmt.Printf("the number of rows updated are %d\n", updateResultCount)
	}

	deleteStmt, err := db.Prepare(`
		DELETE FROM users
		WHERE id = $1
	`)
	if err != nil {
		fmt.Println("There is an error while preparing the delete steatment")
		fmt.Printf("%v\n", err)
	}

	deleteResult, deleteResultErr := deleteStmt.Exec(2)
	if deleteResultErr != nil {
		fmt.Println("An error occured while deleting the row")
		fmt.Printf("%v\n", deleteResultErr)
	}
	deleteResultCount, deleteResultCountErr := deleteResult.RowsAffected()
	if deleteResultCountErr != nil {
		fmt.Println("Error Error Error Error Error")
		fmt.Printf("%v\n", deleteResultCountErr)
	} else {
		fmt.Printf("the number of rows deleted are %d\n", deleteResultCount)
	}

	defer db.Close()
}
