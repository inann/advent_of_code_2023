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
	var gameIDTotal = 0
	redBeadsAvailable := 12
	greenBeadsAvailable := 13
	blueBeadsAvailable := 14

	// Regex for ID
	gameIDRegex := regexp.MustCompile("Game ([0-9]+)")
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
		gameID := gameIDRegex.FindStringSubmatch(initialText)[1]

		_, initialWithoutGameString, _ := strings.Cut(initialText, ":")
		sliceOfGames := strings.Split(initialWithoutGameString, ";")
		validGame := true

		for _, game := range sliceOfGames{
			// For each game, get bead distribution
			// Get slices with individual bead counts
			// fmt.Println(game)
			sliceOfBeads := strings.Split(game, ",")

			for beadNumber, beadDistro := range sliceOfBeads {
				beadCount := gameRegex.FindStringSubmatch(beadDistro)[1]
				beadColor := gameRegex.FindStringSubmatch(beadDistro)[2]
				// fmt.Println("Getting " + beadCount + " for color " + beadColor)
				// for bead type, check if count is less than or equal to allowed quantity
				switch beadColor {
				case "red":
					i, err := strconv.Atoi(beadCount)
					check_error(err)
					if i > redBeadsAvailable {
						fmt.Println("Breaking on Red on bead " + string(beadNumber))
						validGame = false
						break
					}
				case "blue":
					i, err := strconv.Atoi(beadCount)
					check_error(err)
					if i > blueBeadsAvailable {
						fmt.Println("Breaking on Blue on bead " + string(beadNumber))
						validGame = false
						break
					}
				case "green":
					i, err := strconv.Atoi(beadCount)
					check_error(err)
					if i > greenBeadsAvailable {
						fmt.Println("Breaking on Green on bead " + string(beadNumber))
						validGame = false
						break
					}
				}
			}
		}
		if validGame {
			gameIDAsInt, error2 := strconv.Atoi(gameID)
			check_error(error2)
			gameIDTotal = gameIDTotal + gameIDAsInt
		}
	}

	// Print game ID total
	fmt.Println(gameIDTotal)

	// Check scanner for errors
	check_error(scanner.Err())
}