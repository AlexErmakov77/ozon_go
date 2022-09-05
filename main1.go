package main

import "fmt"

func main() {

	// Place your code here.
var str string = "Hello,  Otus!"
strrune := []rune(str)
lenstr := len(strrune)-1


fmt.Print(str, "\n")

fmt.Print(lenstr+1,"\n")

    for  i, j := 0, lenstr ; i < lenstr/2 ; i, j = i+1 , j-1 {

    strrune[i], strrune[j] = strrune[j], strrune[i]
    
    }
strrev :=  string(strrune)
fmt.Print(strrev)

}
