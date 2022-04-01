package main

import "fmt"

type Article struct {
	Title string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

func (a Article) SayHi() string {
	return fmt.Sprintf("hiiii")
}

func (a Article) GetPages() int {
	return 0
}


type Book struct {
	Title string
	Author string
    Pages int
}

func (a Book) String() string {
	return fmt.Sprintf("The %q book was written by %s.", a.Title, a.Author)
}

func (b Book) GetPages() int {
	return b.Pages
}


type Publication interface {
    String() string
    GetPages() int
}

func Print(p Publication) {
    fmt.Println(p.String())
    // fmt.Println("%s has %d pages", p.GetTitle(), p.GetPages())
}


func main() {
	a := Article{
		Title: "my artical",
		Author: "anji",
	}
    b := Book{
        Title: "how to write book",
        Author: "also anji",
    }
	Print(a)
    Print(b)
}








