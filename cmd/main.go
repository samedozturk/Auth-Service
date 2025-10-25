package main

import (
	"fmt"

	"github.com/samedozturk/Auth-Service/internal"
)

func main() {
	internal.StartMongoDB()
	fmt.Println("basarili")
}
