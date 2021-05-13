package heaps

// Permutations uses the Heap's Algorithm to generate all the
// possible permutations of the input.
func Permutations(x []interface{}) [][]interface{} {
	n := len(x)
	res := make([][]interface{}, 0, factorial(n)+1)
	c := make([]int, n)

	res = append(res, copyX(x))

	for i := 1; i < n; {
		if c[i] < i {
			if i%2 == 0 {
				x[0], x[i] = x[i], x[0]
			} else {
				x[c[i]], x[i] = x[i], x[c[i]]
			}

			res = append(res, copyX(x))
			c[i]++
			i = 1
		} else {
			c[i] = 0
			i++
		}
	}

	return res
}

func copyX(x []interface{}) []interface{} {
	res := make([]interface{}, len(x))
	copy(res, x)

	return res
}

func factorial(n int) int {
	res := 1

	for ; n != 0; n-- {
		res *= n
	}

	return res
}
