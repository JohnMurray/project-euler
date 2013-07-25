package main

func main() {
	// since we're deviding from 1..20 evenly, we can start
	// at 40 (could find a better starting point) and increment
	// by 20 each time. We also don't need to check things
	// that divide evenly into 20 (1, 2, 4, 5, 10, 20)
	var num int = 40
	var inc int = 20
	divisors := []int{3, 6, 7, 8, 9, 11, 12, 13, 14, 15,
		16, 17, 18, 19}

	for true {
		all_divisible := true
		for i := 0; i < len(divisors); i++ {
			if num%divisors[i] != 0 {
				all_divisible = false
				break
			}
		}

		if all_divisible {
			println(num)
			return
		} else {
			num += inc
		}
	}
}
