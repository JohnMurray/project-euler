package main

func main() {
	factor_num := 600851475143
	factors := factor_gen(factor_num)

	for factor := factors(); factor > 1; factor = factors() {
		if is_prime(factor) {
			println(factor)
			return
		}
	}
	println(1) // if nothing else is found
}

/*
 * Return a generator instead of calculating them all
 * at once.
 */
func factor_gen(x int) func() int {
	n := x

	return func() int {
		if n&1 == 0 {
			factor := n
			n /= 2
			return factor
		}
		f := 3
		for n >= f*f {
			if n%f == 0 {
				factor := n
				n /= f
				return factor
			} else {
				f++
			}
		}
		if n != 1 {
			factor := n
			n = 1
			return factor
		}

		return 0 // default return
	}
}

/*
 * Just brute force it... I suppose.
 */
func is_prime(x int) bool {
	for i := x - 1; i > 1; i-- {
		if x%i == 0 {
			return false
		}
	}
	return true
}
