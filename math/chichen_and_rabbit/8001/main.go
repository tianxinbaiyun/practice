package main

import "fmt"

func main() {
	chicken, rabbit := ForceSolve(150, 350)
	fmt.Println(chicken, rabbit)
}

//ForceSolve rabbit
func ForceSolve(head, foot int) (chicken, rabbit int) {
	for chicken = 0; chicken < head; chicken++ {
		rabbit = head - chicken
		if 2*chicken+4*rabbit == foot {
			return
		}

	}
	return
}
