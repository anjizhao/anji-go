
package main

import "fmt"

func main() {
    var x string = "helloo world"
    y := x + "!!!!!"
    fmt.Println(x)
    fmt.Println(y)
    y += `:
    P
    `
    fmt.Println(y)
}







