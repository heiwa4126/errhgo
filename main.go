package main

import (
	"fmt"
)

var (
	// Version ...
	Version string = "v9.9.9"
	// Revision =$(git rev-parse --short HEAD)
	Revision string = "9999999"
)

func main() {
	fmt.Printf("errhgo %s (%s)\n", Version, Revision)
}
