// // // // // // // package main

// // // // // // // import (
// // // // // // // 	"database/sql"
// // // // // // // 	"fmt"
// // // // // // // 	_ "fmt"

// // // // // // // 	_ "github.com/lib/pq"
// // // // // // // )

// // // // // // // func main() {
// // // // // // // 	db, err := sql.Open("postgres", "user=ousta password=Password12 host=127.0.0.1 port=5433 dbname=postgres sslmode=disable")

// // // // // // // 	if err != nil {
// // // // // // // 		panic(err)
// // // // // // // 	} else {
// // // // // // // 		fmt.Println("The connection to the DB was successfully initialized!")
// // // // // // // 	}

// // // // // // // 	connectivity := db.Ping()
// // // // // // // 	if connectivity != nil {
// // // // // // // 		panic(err)
// // // // // // // 	} else {
// // // // // // // 		fmt.Println("Good to go!")
// // // // // // // 	}

// // // // // // // 	defer db.Close()
// // // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"database/sql"
// // // // // // 	"fmt"

// // // // // // 	_ "github.com/lib/pq"
// // // // // // )

// // // // // // func main() {
// // // // // // 	db, err := sql.Open("postgres", "user=ousta password=Password12 host=127.0.0.1 port=5433 dbname=postgres sslmode=disable")
// // // // // // 	if err != nil {
// // // // // // 		panic(err)
// // // // // // 	} else {
// // // // // // 		fmt.Println("The connection to the DB was successfully initialized!")
// // // // // // 	}

// // // // // // 	// DBCreate := `
// // // // // // 	// CREATE TABLE public.test2
// // // // // // 	// (
// // // // // // 	// 	id integer PRIMARY KEY,
// // // // // // 	// 	name character varying COLLATE pg_catalog."default"
// // // // // // 	// )
// // // // // // 	// WITH (
// // // // // // 	// 	OIDS = FALSE
// // // // // // 	// )
// // // // // // 	// TABLESPACE pg_default;
// // // // // // 	// ALTER TABLE public.test2
// // // // // // 	// 	OWNER to postgres;
// // // // // // 	// `
// // // // // // 	// _, err = db.Exec(DBCreate)
// // // // // // 	if err != nil {
// // // // // // 		panic(err)
// // // // // // 	} else {
// // // // // // 		fmt.Println("The table was successfully created!")
// // // // // // 	}

// // // // // // 	insert, err := db.Prepare("INSERT INTO test2(id, name) VALUES ($1, $2)")

// // // // // // 	if err != nil {
// // // // // // 		panic(err)
// // // // // // 	}

// // // // // // 	_, err = insert.Exec(3, "third")
// // // // // // 	if err != nil {
// // // // // // 		panic(err)
// // // // // // 	}
// // // // // // 	fmt.Println("The value was successfully inserted!")
// // // // // // 	defer db.Close()
// // // // // // }

// // // // // package main

// // // // // import (
// // // // // 	"database/sql"
// // // // // 	"fmt"

// // // // // 	_ "github.com/lib/pq"
// // // // // )

// // // // // func main() {
// // // // // 	var property string

// // // // // 	db, err := sql.Open("postgres", "user=ousta password=Password12 host=127.0.0.1 port=5433 dbname=postgres sslmode=disable")
// // // // // 	if err != nil {
// // // // // 		panic(err)
// // // // // 	} else {
// // // // // 		fmt.Println("The connection to the DB was successfully initialized!")
// // // // // 	}

// // // // // 	TableCreate := `
// // // // // 	CREATE TABLE Number
// // // // // 	(
// // // // // 	  Number integer NOT NULL,
// // // // // 	  Property text COLLATE pg_catalog."default" NOT NULL
// // // // // 	)
// // // // // 	WITH (
// // // // // 	  OIDS = FALSE
// // // // // 	  )
// // // // // 	TABLESPACE pg_default;
// // // // // `

// // // // // 	_, err = db.Exec(TableCreate)
// // // // // 	if err != nil {
// // // // // 		panic(err)
// // // // // 	} else {
// // // // // 		fmt.Println("The table called Numbers was successfully created!")
// // // // // 	}

// // // // // 	insert, insertErr := db.Prepare("INSERT INTO Number VALUES($1,$2)")
// // // // // 	if insertErr != nil {
// // // // // 		panic(insertErr)
// // // // // 	}
// // // // // 	for i := range 100 {
// // // // // 		if i%2 == 0 {
// // // // // 			property = "Even"
// // // // // 		} else {
// // // // // 			property = "Odd"
// // // // // 		}
// // // // // 		_, err = insert.Exec(i, property)
// // // // // 		if err != nil {
// // // // // 			panic(err)
// // // // // 		} else {
// // // // // 			fmt.Println("The number:", i, "is:", property)
// // // // // 		}
// // // // // 	}
// // // // // 	insert.Close()
// // // // // 	fmt.Println("The numbers are ready.")

// // // // // 	db.Close()
// // // // // }

// // // // package main

// // // // import (
// // // // 	"database/sql"
// // // // 	"fmt"

// // // // 	_ "github.com/lib/pq"
// // // // )

// // // // func main() {

// // // // 	var number int
// // // // 	var property string

// // // // 	db, err := sql.Open("postgres", "user=ousta password=Password12 host=127.0.0.1 port=5433 dbname=postgres sslmode=disable")

// // // // 	if err != nil {
// // // // 		panic(err)
// // // // 	} else {
// // // // 		fmt.Println("The connection to the DB was successfully initialized!")
// // // // 	}

// // // // 	rows, err := db.Query("SELECT * FROM number")
// // // // 	if err != nil {
// // // // 		panic(err)
// // // // 	}

// // // // 	for rows.Next() {
// // // // 		err := rows.Scan(&number, &property)
// // // // 		if err != nil {
// // // // 			panic(err)
// // // // 		}

// // // // 		fmt.Printf("Retrieved data from db: %d %s\n", number, property)
// // // // 	}

// // // // 	err = rows.Err()
// // // // 	if err != nil {
// // // // 		panic(err)
// // // // 	}

// // // // 	err = rows.Close()
// // // // 	if err != nil {
// // // // 		panic(err)
// // // // 	}
// // // // 	db.Close()
// // // // }

// // // package main

// // // import (
// // // 	"database/sql"
// // // 	"fmt"

// // // 	_ "github.com/lib/pq"
// // // )

// // // func main() {
// // // 	var name string
// // // 	id := 2

// // // 	db, err := sql.Open("postgres", "user=ousta password=Password12 host=127.0.0.1 port=5433 dbname=postgres sslmode=disable")
// // // 	if err != nil {
// // // 		panic(err)
// // // 	} else {
// // // 		fmt.Println("The connection to the DB was successfully initialized!")
// // // 	}
// // // 	qryrow, err := db.Prepare("SELECT name FROM test WHERE id=$1")
// // // 	if err != nil {
// // // 		panic(err)
// // // 	}
// // // 	err = qryrow.QueryRow(id).Scan(&name)
// // // 	if err != nil {
// // // 		panic(err)
// // // 	}
// // // 	fmt.Printf("The name with id %d is %s\n", id, name)
// // // 	err = qryrow.Close()
// // // 	if err != nil {
// // // 		panic(err)
// // // 	}
// // // 	db.Close()
// // // }

// // package main

// // import (
// // 	"database/sql"
// // 	"fmt"

// // 	_ "github.com/lib/pq"
// // )

// // func main() {
// // 	db, err := sql.Open("postgres", "user=ousta password=Password12 host=127.0.0.1 port=5433 dbname=postgres sslmode=disable")
// // 	if err != nil {
// // 		panic(err)
// // 	} else {
// // 		fmt.Println("The connection to the DB was successfully initialized!")
// // 	}

// // 	updateStatement := `
// // 	UPDATE test
// // 	SET name = $1
// // 	WHERE id = $2
// // 	`
// // 	updateResult, updateResultErr := db.Exec(updateStatement, "well", 22)
// // 	if updateResultErr != nil {
// // 		panic(updateResultErr)
// // 	}
// // 	updatedRecords, updatedRecordsErr := updateResult.RowsAffected()
// // 	if updatedRecordsErr != nil {
// // 		panic(updatedRecordsErr)
// // 	}
// // 	fmt.Println("Number of records updated: ", updatedRecords)
// // 	db.Close()
// // }

// package main

// import (
// 	"database/sql"
// 	"flag"
// 	"fmt"
// 	"os"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	var id int
// 	flag.IntVar(&id, "id", 0, "The id to be deleted")
// 	flag.Parse()

// 	if id == 0 {
// 		fmt.Println("The id arg is required")
// 		os.Exit(1)
// 	}

// 	db, err := sql.Open("postgres", "user=ousta password=Password12 host=127.0.0.1 port=5433 dbname=postgres sslmode=disable")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Connection to the database has been successfully")

// 	deleteSteatment := `
// 		DELETE FROM test
// 		WHERE ID = $1
// 	`
// 	deleteRow, deleteRowErr := db.Exec(deleteSteatment, id)
// 	if deleteRowErr != nil {
// 		panic(deleteRowErr)
// 	}

// 	deltedRowsCount, deltedRowsCountErr := deleteRow.RowsAffected()
// 	if deltedRowsCountErr != nil {
// 		panic(deltedRowsCountErr)
// 	}

// 	fmt.Printf("The numbers of rows deleted are: %d\n", deltedRowsCount)

// 	defer db.Close()

// }
package main
