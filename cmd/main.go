package main

import (
	"fmt"
	parser "github.com/codescalersinternships/dotenv-Fatma-Ebrahim/pkg"
)

func main(){
	// res, err := parser.ParseString("KEY=VALUE\n        KEY2=      \"VALUE2\n value with spaces\"\nKEY3=VALUE3")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Parsed Result:", res["KEY2"])

	res, err := parser.ParseFile("./.env")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for key, value := range res {
		fmt.Printf("%s=%s\n", key, value)
	}
	fmt.Println("Parsed Result:", res)

}

// 
// export KEY3=VALUE3 # with spaces
// export KEY4="VALUE WITH SPACES"
// # with multi lines

// KEY6="Value with a # hash"