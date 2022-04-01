
package main

import "fmt"

func main1() {
    var x [5]int
    fmt.Println(x)
    x[4] = 100
    fmt.Println(x)
}


func avg(x []float64) float64 {
    var sum float64 = 0
    for _, val := range x {
        sum += val
    }
    return sum / float64(len(x))
}

func main() {
    x := [5]float64{
        98, 93, 77, 82, 83,
    }
    fmt.Printf("x: %T\n", x)
    x_slice := x[:]
    fmt.Printf("x_slice: %T\n", x_slice)
    result := avg(x_slice)
    fmt.Printf("result: %T\n", result)
    fmt.Println(result)
}






