package dna

import (
	"fmt"
)

type Histogram map[byte]int

type Dna map[byte]int

func DNA(s string) Dna {
	d := Dna{'G': 0, 'T': 0, 'A': 0, 'C': 0}
	for _, n := range s {
		d[byte(n)]++
	}
	return d
}

func (d Dna) Count(n byte) (int, error) {
	count, ok := d[n]
	if !ok {
		return 0, fmt.Errorf("%q invalid nucleotide", n)
	}
	return count, nil
}

func (d Dna) Counts() Histogram {
	return Histogram(d)
}
