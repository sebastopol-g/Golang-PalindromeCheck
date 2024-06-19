package main

import (
	"fmt"
	"bufio"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func Reverse(s string) (result string) {
	for _,v := range s {
	  result = string(v) + result
	}
	return 
}

func IsPalindrome(word string) bool {
	size := len(word)
	var isPal bool

	if size != 0 {
		reversedWord := Reverse(word)
		if reversedWord == word {
			isPal = true
		}
	} else {
		fmt.Print("Empty word")
		isPal = false
	}

	return isPal
}

func main() {
		scanner := bufio.NewScanner(os.Stdin)
    	var text string
   		for text != "q" { 
			fmt.Print("Enter your text: ")
			scanner.Scan()
			text = scanner.Text()
			if text != "q" {
				fmt.Printf("Your text was %v\n", text)
				result := IsPalindrome(text)
				if result == true {
					fmt.Printf("%v is a palindrome :)\n", text)
				} else {
					fmt.Printf("%v is NOT a palindrome :(\n", text)
				}
			}
		}
}