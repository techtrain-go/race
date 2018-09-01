package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/techtrain-go/race/solver"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	rawData, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	start, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	dest, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	// Code for getting inputs from os.Args[1]
	_, err = solver.Solve(points, start, dest)
	if err != nil {
		log.Fatal(err)
	}
}
