
package main

import (
    "fmt"
    "strconv"
    "time"
)

// Declare a Book type which satisfies the fmt.Stringer interface.
type Book struct {
    Title  string
    Author string
}

func (b Book) String() string {
    return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

// Declare a Count type which satisfies the fmt.Stringer interface.
type Count int

func (c Count) String() string {
    return strconv.Itoa(int(c))
}


// Declare a WriteLog() function which takes any object that satisfies
// the fmt.Stringer interface as a parameter.
func WriteLog(s fmt.Stringer) {
    fmt.Printf("%s | %s\n", time.Now(), s.String())
    // fmt.Println(s + 1)
    // s now has type Stringer and thus does not support int operations like +
    // we can use type assertion to check whether the value passed into this function
    // has static type Count (which has underlying type int which is addable)
    c, ok := s.(Count)
    if ok {
        fmt.Printf("%s | %s\n", time.Now(), c + 1)
    } else {
        fmt.Printf("error: cannot add to %T type\n", s)
    }
}

func main() {
    // Initialize a Book object and pass it to WriteLog().
    book := Book{"Alice in Wonderland", "Lewis Carrol"}
    WriteLog(book)

    // Initialize a Count object and pass it to WriteLog().
    count := Count(3)
    WriteLog(count)
}












