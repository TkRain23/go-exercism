// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	Equ Kind = iota // equilateral
	NaT             // not a triangle
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides should have a comment documenting it.
const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	} else if isEquilateral(a, b, c) {
		return Equ
	} else if isIsoceles(a, b, c) {
		return Iso
	} else if isScalene(a, b, c) {
		return Sca
	}
	return NaT
}

func isTriangle(a, b, c float64) bool {
	return isSide(a) && isSide(b) && isSide(c) &&
		a+b >= c && a+c >= b && b+c >= a
}

func isSide(f float64) bool {
	return f > 0 && !math.IsNaN(f) && !math.IsInf(f, 1)
}

func isEquilateral(a, b, c float64) bool {
	return a == b && a == c && b == c
}

func isIsoceles(a, b, c float64) bool {
	return a == b || b == c || a == c
}

func isScalene(a, b, c float64) bool {
	return a != b && a != c && b != c
}
