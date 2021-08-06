package main
import "fmt"


func main() {
  switch 6 {
  case 1,5,10:
	fmt.Println("one")
  case 2,4,6:
	fmt.Println("two")
  default:
	fmt.Println("another number")
  }
}