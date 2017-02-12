// EulerGo is a command line utility to manage solutions in go to the Project Euler Challenges.
package main

import (
	"fmt"

	euler "github.com/Nooby/EulerGo/lib"
)

func main() {
	v1, err := euler.Verify(1, "233168")
	if err != nil {
		fmt.Println(fmt.Errorf("verify failed: %v", err))
	}
	fmt.Println(v1)

	v2, err := euler.Verify(1, "233260")
	if err != nil {
		fmt.Println(fmt.Errorf("verify failed: %v", err))
	}
	fmt.Println(v2)

	//fmt.Println(ep.descriptions[174])
	//fmt.Println(ep.solutions[0])
	//fmt.Println(ep.solutions[463])

}
