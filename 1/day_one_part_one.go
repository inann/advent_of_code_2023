package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)


func check_error(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	fmt.Println("Hello World")

	// Get arguments
	allArgs := os.Args

	// Define Globals
	// Get regex
	findDigits := regexp.MustCompile("\\d")
	var runningTotal = 0

	// First argument is program name
	// Second argument is file name (by convention)
	// Error check for length of args
	if len(allArgs) < 2 {
		os.Exit(1)
	}

	// Grab it from all args
	filename := allArgs[1]

	// Read using ReadFile?
	fileHandle, err := os.Open(filename)

	// Check for error
	check_error(err)

	// Assuming we're here, we have valid open file
	// Defer for closing
	defer fileHandle.Close()

	// Get a Scanner object
	scanner := bufio.NewScanner(fileHandle)

	// Assume no lines over 2^16 in length
	// Get lines
	for scanner.Scan() {
		// For each line, regex check for numbers
		digits := findDigits.FindAllString(scanner.Text(), -1)
		
		// Acquired Numbers, get first and last and make "one" number

		firstDigit := digits[0]
		lastDigit := digits[len(digits) - 1]
		lineValue := firstDigit + lastDigit

		// Parse and check for errors
		intValue, err := strconv.Atoi(lineValue)

		check_error(err)

		// No errors in parse, add to running total
		runningTotal = runningTotal + intValue
	}

	fmt.Println(runningTotal)
	// Check scanner for errors
	check_error(scanner.Err())
}
