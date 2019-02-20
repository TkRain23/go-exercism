package hamming

import "errors"

func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("Hamming distance is only defined for sequences of equal length.")
	}

	hamDist := 0

	for i := range a {
		if a[i] != b[i] {
			hamDist++
		}
	}

	return hamDist, nil

}
