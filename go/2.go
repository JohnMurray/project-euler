package main

func main() {
	var sum int64 = 0
	var max int64 = 4000000
	fib := fib_generator()

	for i := fib(); i < max; i = fib() {
		if i%2 == 0 {
			sum += i
		}
	}

	println(sum)
}

func fib_generator() func() int64 {
	var prev int64 = 0
	var curr int64 = 1
	return func() int64 {
		temp := curr
		curr = curr + prev
		prev = temp
		return curr
	}
}
