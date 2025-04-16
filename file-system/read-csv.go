package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	in := `firstname,lastname,fullname,age
Oussama,Tailba,Oussama Tailba,25
Kaoutar,Taki,kaoutar Taki,22
`

	r := csv.NewReader(strings.NewReader(in))

	for {
		row, err := r.Read()

		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		fmt.Println(row[2])

	}

	os.Remove("temp.txt")

}
