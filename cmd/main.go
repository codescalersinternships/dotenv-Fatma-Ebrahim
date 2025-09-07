package main

import (
	"fmt"
	parser "github.com/codescalersinternships/dotenv-Fatma-Ebrahim/pkg"
)

func main(){
	res, err := parser.ParseString("KEY=VALUE\n        KEY2=      VALUE2\n value with spaces\"\nKEY3=VALUE3")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Parsed Result:", res["KEY3"])

}