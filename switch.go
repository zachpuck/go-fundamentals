//switch
// default statement can go anywhere
// switch and case types must match

package main

import (
	"fmt"
)

func main() {
	switch "docker" {
	case "go":
		fmt.Println("\nHere are some recommended" +
			" go courses...")
	case "kubernetes":
		fmt.Println("\nHere are some recommended" +
			" kubernetes courses...")
	case "docker":
		fmt.Println("\nHere are some recommended" +
			" docker courses...")
	default:
		fmt.Println("\nNot sure where to start?" +
			" check out javascript")
	}
}