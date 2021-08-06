package main
import "fmt"

type Doctor struct {
	number int
	actor_name string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number: 3,
		actor_name: "Zendaya",
		companions: []string {
			"Daximus Henry",
			"Dax Moody",
			"Dot Benji",
},

}
fmt.Println(aDoctor)
}