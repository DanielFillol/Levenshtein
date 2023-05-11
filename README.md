# Levenshtein Distance Calculation
This is a Go package that provides a function to calculate the [Levenshtein distance](https://en.wikipedia.org/wiki/Levenshtein_distance) between two strings,the distance is defined as the minimum number of edits needed to transform one string into the other, with the allowable edit operations being insertion, deletion, or substitution of a single character

This implementation is optimized to use O(min(m,n)) space.It is based on the optimized C version found [here](http://en.wikibooks.org/wiki/Algorithm_implementation/Strings/Levenshtein_distance#C)


## Install
To install this package, use go get:
``` 
go get github.com/Darklabel91/Levenshtein
``` 

## Usage
To use this package in your Go program, import it and call the Distance function with two strings:
```go
package main

import (
	"fmt"
	"github.com/Darklabel91/Levenshtein"
)

func main() {
    str1 := "kitten"
    str2 := "sitting"
    dist := levenshtein.Distance(str1, str2)
    fmt.Printf("The Levenshtein distance between %q and %q is %d\n", str1, str2, dist)
}
``` 
This will be the output:
``` 
The Levenshtein distance between "kitten" and "sitting" is 3
``` 

## Testing
This package includes a test file levenshtein_test.go that tests the Distance function with various inputs. To run the tests, use go test:
```go
go test
```

## Contributing
Contributions are welcome! If you find a bug or would like to add a feature, please open an issue or pull request.
