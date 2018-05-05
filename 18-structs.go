// custom data types
package main

import (
    "fmt"
)

func main() {

    type cat struct {
        name string
        color string
        attackRating float64
    }

    // ways to declare and intialize variables of type cat
    var milo cat
    milo.name = "Milo"
    milo.color = "Orange"
    milo.attackRating = 8.7
    
    eva := new(cat) // this one gives us a pointer
    eva = &cat{ "Eva", "Gray", 3.6}
    
    oreo := cat{
        name: "Oreo",
        color: "black and white",
        attackRating: 7.2,
    }

    fmt.Printf("milo: %v,\n eva: %v,\n oreo: %v", milo, *eva, oreo)

    fmt.Println("\nMilo has the highest attack rating of", milo.attackRating)
}