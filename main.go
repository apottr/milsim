package main

import (
	"fmt"

	t "github.com/apottr/milsim/grid"
	u "github.com/apottr/milsim/unit"
)

func main() {
	fmt.Println("Hello World!")
	group := []*u.Force{
		u.CreateForce("MyForce", 5000),
		u.CreateForce("SecondForce", 100),
		u.CreateForce("thirdForce", 50),
		u.CreateForce("FourthForce", 10000),
	}
	a := t.CreateTile("UTM11", group)
	a.Pp()

}
