package main

import (
	"fmt"
)

//Euclid algorithm, time O(log(B))
func euclid_algorithm() {
	A := 4851
	B := 3003

	//Just save initial numbers
	a := A
	b := B

	for B != 0 {
		//First we must found reminder
		var rem int = (A) % (B)

		//Then we change numbers (trying to found minimum number)
		A = B
		B = rem
	}

	//GCD(A, B) = GCD(B, A mod B), if R = A mod B, then A = m x B + R, where m some int. (it's just my notes)
	//If GCD for 2 numbers are 1, then numbers calls as Mutually prime
	if A == 1 {
		fmt.Print("Numbers are Mutually-prime")
	} else {
		fmt.Printf("Greatest common divisor for integers %d and %d is %d", a, b, A)
	}
}
