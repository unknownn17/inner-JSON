package main

import (
	"encoding/json"
	"fmt"
	//"log"
	"os"
)

/* Load the JSON file into your program.
Print the title and author of each book.
Count the number of books in the file.
Find and print the books published before 1950.
Add a new book to the JSON file and save it.*/
type Books struct {
    Title string `json:"title"`
    Author  string   `json:"author"`
	Genre  string   `json:"genre"`
	Year  int   `json:"year"`

}
type All struct{
	Books []Books `json:"book"`
}

func encoder(){
    file, _ := os.Create("input.json")
    defer file.Close()

    book := All{
		[]Books{
		{Title: "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Genre: "Fiction",
		Year: 1925},		  
		{Title: "To Kill a Mockingbird",
		Author: "Harper Lee",
		Genre: "Fiction",
		Year: 1960},
		{Title: "1984",
		Author: "George Orwell",
		Genre: "Science Fiction",
		Year: 1945},
	},
}
	_=json.NewEncoder(file).Encode(book)
}
    



func decoder(){
 file, _ := os.Open("input.json")
 defer file.Close()

 var book All
 _=json.NewDecoder(file).Decode(&book)

fmt.Println("--------- Title and Author of each books ----------")

 for _,t:=range book.Books{
	fmt.Printf("Title: %v\nAuthor: %v\n",t.Title,t.Author)
	fmt.Print("\n")
 }
 fmt.Println("--------- Books that Published before 1950 ----------")
 for _,t:=range book.Books{
	if t.Year<1950{
		fmt.Println("Title: ",t.Title)
		fmt.Println("Author: ",t.Author)
		fmt.Println("Genre: ",t.Genre)
		fmt.Println("Year: ",t.Year)
		fmt.Print("\n")
	}
 }

 b:=Books{Title: "Animal Farm",Author: "George Orwell",Genre: "Adventure",Year: 1980}

 book.Books=append(book.Books, b)
 fmt.Println("Total Books: ",len(book.Books))
}

func main() {
   encoder()
   decoder()
}
