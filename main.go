package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	// Version ...
	Version string = "v9.9.9"
	// Revision =$(git rev-parse --short HEAD)
	Revision string = "9999999"
)

func main() {
	fmt.Printf("errhgo %s (%s)\n", Version, Revision)

	read0("testdata/main/0.txt")
	read0("testdata/main/1.txt")
	read0("testdata/main/2.txt")
	read0("testdata/main/3.txt")
}

func read0(filename string) {
	b, err := readTrue(filename)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("%s: %t\n", filename, b)
	}
}

func readTrue(filename string) (bool, error) {

	f, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer f.Close()

	r := bufio.NewScanner(f)
	if r.Scan() {
		switch r.Text() {
		case "true":
			return true, nil
		case "false":
			return false, nil
		}
	}

	if err = r.Err(); err == nil {
		err = fmt.Errorf("not sure")
	}
	return false, err
}
