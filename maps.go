package main
import "fmt"


func main() {
  statesPopulations := map[string]int{
	"California" : 3950,
	"Texas" : 4950,
	"Florida" : 2069,
}
statesPopulations["Georgia"]=9087
fmt.Println(statesPopulations)
}