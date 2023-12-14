package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func check_error(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	// Get arguments
	allArgs := os.Args

	// Define Globals
	filename := allArgs[1]

	// Regex for ID
	gameRegex := regexp.MustCompile("([0-9]+) (red|blue|green)")

	// Read using Readfile
	fileHandle, err := os.Open(filename)

	// check for error
	check_error(err)

	// Assuming we're here, we have valid open file
	// Defer for closing
	defer fileHandle.Close()

	// Get a Scanner object
	scanner := bufio.NewScanner(fileHandle)
	allGamesPowersSum := 0

	// Assume no lines over 2^16 in length
	// Get lines
	for scanner.Scan() {
		// For each line, gather information from the text
		initialText := scanner.Text()

		// Info to gather:
		//	- Game ID
		//	- Split On Games
		//	- Info On Each Game
		//	- Check On If Game Is Valid
		// gameID := gameIDRegex.FindStringSubmatch(initialText)[1]
		gamePower := 0

		_, initialWithoutGameString, _ := strings.Cut(initialText, ":")
		sliceOfDraws := strings.Split(initialWithoutGameString, ";")

		// REDO FOR NEW PART
		redMin, blueMin, greenMin := 0, 0, 0
		// validGame := true

		for _, draw := range sliceOfDraws{
			// For each game, get bead distribution
			// Get slices with individual bead counts
			// fmt.Println(game)
			sliceOfBeads := strings.Split(draw, ",")

			for _, beadDistro := range sliceOfBeads {
				beadCount := gameRegex.FindStringSubmatch(beadDistro)[1]
				beadColor := gameRegex.FindStringSubmatch(beadDistro)[2]
				// fmt.Println("Getting " + beadCount + " for color " + beadColor)
				// for bead type, check if count is less than or equal to allowed quantity
				switch beadColor {
				case "red":
					i, err := strconv.Atoi(beadCount)
					check_error(err)
					if i > redMin {
						// REDO FOR NEW PART
						// validGame = false
						redMin = i
						break
					}
				case "blue":
					i, err := strconv.Atoi(beadCount)
					check_error(err)
					if i > blueMin {
						// REDO FOR NEW PART
						// validGame = false
						blueMin = i
						break
					}
				case "green":
					i, err := strconv.Atoi(beadCount)
					check_error(err)
					if i > greenMin {
						// REDO FOR NEW PART
						// validGame = false
						greenMin = i
						break
					}
				}
			}
		}
		// REDO FOR NEW PART
		// if validGame {
		// 	gameIDAsInt, error2 := strconv.Atoi(gameID)
		// 	check_error(error2)
		// 	gameIDTotal = gameIDTotal + gameIDAsInt
		// }
		gamePower = redMin * blueMin * greenMin
		allGamesPowersSum = allGamesPowersSum + gamePower
	}

	// Print game ID total
	fmt.Println(allGamesPowersSum)

	// Check scanner for errors
	check_error(scanner.Err())
}