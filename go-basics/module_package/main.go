package main

import (
	"fmt"
	"os"
	"module_package/calucator"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("GO_ENV"))

	fmt.Println(calucator.Offset)
	fmt.Println(calucator.Sum(1, 2))
	fmt.Println(calucator.Multiply(1, 2))
}