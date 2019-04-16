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