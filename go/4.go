package main

func main() {
	var largest_palindrome int = 0

	for i := 999; i > 99; i-- {
		for j := i; j > 99; j-- {
			if is_palindrome(i*j) && (i*j) > largest_palindrome {
				largest_palindrome = (i * j)
			}
		}
	}

	println(largest_palindrome)
}

func is_palindrome(num int) bool {
	rev := 0
	n := num

	for n > 0 {
		rev = rev*10 + (n % 10)
		n /= 10
	}

	return num == rev
}
