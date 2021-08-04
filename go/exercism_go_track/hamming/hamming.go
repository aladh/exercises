package hamming

import "errors"

// Distance calculates the hamming distance between a and b
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("both strands should have the same length")
	}

	hammingDistance := 0

	for i := range a {
		if a[i] != b[i] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
