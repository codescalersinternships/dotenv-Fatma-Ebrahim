package main

import (
	"fmt"
	"os"

	parser "github.com/codescalersinternships/dotenv-Fatma-Ebrahim/pkg"
)

func main(){
	res, err := parser.ParseString("KEY=VALUE\n        KEY2=      \"VALUE2\n value with spaces\"\nKEY3=VALUE3")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Parsed Result of ParseString:", res)

	res, err = parser.ParseFile("./.env")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Parsed Result of ParseFile:", res)

	err=parser.LoadEnvFile("./.env")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(os.Getenv("KEY"))
	os.Unsetenv("KEY")
}

