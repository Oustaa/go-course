package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://google.com",
		"Facebook": "https://facebook.com",
		"Amazon web services": "https://aws.com",
	}

	// fmt.Println(websites)
	// websites["LinkedIn"] = "https://linkedin.com"
	// delete(websites, "LinkedIn")
	// fmt.Println(websites)

	for key, value := range websites {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
		fmt.Println("---------------")
	}

}