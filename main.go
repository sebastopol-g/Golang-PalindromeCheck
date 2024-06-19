package main

import (
	"bufio"
	"fmt"
	"os"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func IsPalindrome(word string) bool {
	var isPal bool

	reversedWord := Reverse(word)
	if reversedWord == word {
		isPal = true
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
			if len(text) == 0 {
				fmt.Println("Your word is empty, analyse can't be done")
			} else {
				result := IsPalindrome(text)
				if result {
					fmt.Printf("%v is a palindrome :)\n", text)
				} else {
					fmt.Printf("%v is NOT a palindrome :(\n", text)
				}
			}
		}
	}
}
