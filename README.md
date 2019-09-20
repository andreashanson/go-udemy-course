# Run go program ( This will compile and run the program. )
go run <filename.go>

# Build go program ( This will compile the program. )
go build

# Formats all the code in current project.
go fmt

# Install go program ( This will build and put the file in go/bin directory. )
go install

# Download package
go get <package url>

# Run test for current project
go test


# import
package main
Example import "fmt"

// This gives the package main access to all the code inside packge fmt.
//We have imported the code in package fmt to our package main.

main package is called main.

Other standard packages our.
debug
math
fmt
encoding
crypto
io

To get access to hem in your main package you have to import them.

# Documentation about standard packages like "fmt"
golang.org/pkg

# How is main.go organized?

package main
import "<packagename>"

func main() {
    # Your code here.
}

# Arrays and Slices
An array is a fixed length of a list
A Slice is a list that can grow or shrink

# Slice
 - Add card to slice
  slice = append(cards, "Six of Spades")

# Pointers

func (pointerToPerson *person) updateName() {
    *pointerToPerson
}

Star in front of a type we are looking for a type of pointer to a person

But a star in front of an actual pointer we are looking for the value of that pointer.

Vi behöver inte bry oss om pekare på referens typer.

Vi behöver bry oss om pekare på värde typer.

En slice är en referens typ för den innehåller en referens till en underliggande lista of records.

När vi skall passa in argument till funktioner så försöker vi använda pekaren till värdet för att det inte skall bli fel.

# Maps

This means that we are declaring a map colors.
And we say that all of the keys needs to be of type string.
And all of its values also needs to be of type string.

colors := map[string]string{}