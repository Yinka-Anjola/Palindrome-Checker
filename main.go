package main 

import (
	"fmt"
	"strings"
        "unicode"
)


func main() {
	var input string

	fmt.Println("Enter a word: ")
	_, err := fmt.Scanln(&input)
	if err != nil{
		fmt.Println("Expected string input")
                return
	}
        if !isAlphabetic(input){
                fmt.Println("Input should only contain alphabetic characters")
                return
        }

	if checkIsPalindrome(input) {
		fmt.Println("The word provided is a palindrome")
	} else {
		fmt.Println("The word provided is not a palindrome")
	}
}
func isAlphabetic (str string) bool{
        for _, char := range str{
                if !unicode.IsLetter(char) {
                        return false
                }
        }
        return true
}
func checkIsPalindrome(str string) bool {
	str = strings.ToLower(strings.ReplaceAll(str, " ", ""))
	return str == reverseString(str)
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
