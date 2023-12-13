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
	findDigitsForward := regexp.MustCompile("[0-9]|one|two|three|four|five|six|seven|eight|nine")
	findDigitsBackward := regexp.MustCompile("[0-9]|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin")
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
		// Reverse String
		forwardString := scanner.Text()
		var reverseString = ""
		runeArray := []rune(forwardString)
		for i,j := 0,len(forwardString)-1; i < j; i,j = i+1, j-1 {
			runeArray[i], runeArray[j] = runeArray[j], runeArray[i]
		}
		
		reverseString = string(runeArray)

		// Acquired Numbers, get first and last and make "one" number

		firstDigit := findDigitsForward.FindAllString(forwardString, -1)[0]
		lastDigit := findDigitsBackward.FindAllString(reverseString, -1)[0]
		// Check both first and last and translate from string to different string
		switch firstDigit {
		case "one":
			// fmt.Println(firstDigit)
			firstDigit = "1"
		case "two":
			// fmt.Println(firstDigit)
			firstDigit = "2"
		case "three":
			// fmt.Println(firstDigit)
			firstDigit = "3"
		case "four":
			// fmt.Println(firstDigit)
			firstDigit = "4"
		case "five":
			// fmt.Println(firstDigit)
			firstDigit = "5"
		case "six":
			// fmt.Println(firstDigit)
			firstDigit = "6"
		case "seven":
			// fmt.Println(firstDigit)
			firstDigit = "7"
		case "eight":
			// fmt.Println(firstDigit)
			firstDigit = "8"
		case "nine":
			// fmt.Println(firstDigit)
			firstDigit = "9"
		default:
			// fmt.Println(firstDigit)
			firstDigit = firstDigit
		}

		switch lastDigit {
		case "eno":
			// fmt.Println(lastDigit)
			lastDigit = "1"
		case "owt":
			lastDigit = "2"
		case "eerht":
			lastDigit = "3"
		case "ruof":
			lastDigit = "4"
		case "evif":
			lastDigit = "5"
		case "xis":
			lastDigit = "6"
		case "neves":
			lastDigit = "7"
		case "thgie":
			lastDigit = "8"
		case "enin":
			lastDigit = "9"
		default:
			lastDigit = lastDigit
		}

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
