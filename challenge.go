package main
import (
	"fmt"
	
)

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {

func checkPalindrome(inputString string) bool {
  
	strRev := Reverse(inputString)
	
	if inputString == strRev{
		print("yes")
		return true
	}

	return false
}
}