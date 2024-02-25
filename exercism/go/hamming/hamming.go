package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("length not equal")
	}
	

	hammingDistance := 0
	for pos := range a {
		if a[pos] != b[pos] {
			hammingDistance += 1
		}
	}
	return hammingDistance, nil;
}
