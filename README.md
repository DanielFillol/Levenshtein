# [Levenshtein Distance](http://en.wikipedia.org/wiki/Levenshtein_distance)

The levenshtein distance between two strings is defined as the minimum number of edits needed to transform one string into the other, with the allowable edit operations being insertion, deletion, or substitution of a single character

This implementation is optimized to use O(min(m,n)) space.
It is based on the optimized C version found [here](http://en.wikibooks.org/wiki/Algorithm_implementation/Strings/Levenshtein_distance#C)


## Install
``` 
go get github.com/Darklabel91/Levenshtein
``` 

## Example
```go
package main

import (
	"fmt"
	"github.com/Darklabel91/Levenshtein"
)

func main() {
	levenshtein := Levenshtein.Distance("Daniel", "Danilo")
	fmt.Println(levenshtein)
}

``` 
Return
``` 
2
``` 
