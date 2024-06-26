package main

import (
	"fmt"

	"../pkg/api"
)

func main() {
	fmt.Println("Started the API server")
	api.Start()
}
