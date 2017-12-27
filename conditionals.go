package main

import (
	"fmt"
)

func main() {
	//Variables to store source rankings
	// firstRank := "39"
	// secondRank := "614"

	if firstRank, secondRank := 39, 614; firstRank < secondRank {
		fmt.Println("\nFirst course is doing better" +
			" than second course")
	} else if firstRank > secondRank {
		fmt.Println("\nOh dear.... your first" +
			"course must be doing abysmally!")
	} else {
		fmt.Println("\nBoth courses are either" +
			"the same or something weird is going on")
	}

	// switch
	
}