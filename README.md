# dotenv-Parser
This repository contains an implementation of a Go library that can parse and load environment variables from .env files. This allows for easier configuration management across different environments.

## Functions
### `ParseString(s string) (map[string]string, error)`
takes a string of .env data and returns a map of key-value pairs.

### `ParseFile(path string) (map[string]string, error)`
takes a path to a .env file and returns a map of key-value pairs.

### `LoadEnvString(s string) error`
takes a string of .env data and loads the environment variables into the process.

### `LoadEnvFile(path string) error`
takes a path to a .env file and loads the environment variables into the process.




## How to Use:
### Step 1: Install the library using `go get`

  ```bash
  go get github.com/codescalersinternships/dotenv-Fatma-Ebrahim
  ```

This command fetches the library and adds it to your project's `go.mod` file.

### Step 2: Import and use the library in your code

  After running `go get`, you can import the library into your project and use the functions as described:

```
package main

import (
	"fmt"
	parser "github.com/codescalersinternships/dotenv-Fatma-Ebrahim/pkg"
)

func main(){
	res, err := parser.ParseString("KEY=VALUE")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Parsed Result of ParseString:", res)

	err=parser.LoadEnvFile("./.env")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
```