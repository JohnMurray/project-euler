package main

func main() {
	println(first_100_nat_squared() - first_100_squared())
}

func first_100_squared() int64 {
	var sum int64 = 0
	for i := 1; i <= 100; i++ {
		sum += int64(i * i)
	}

	return sum
}

func first_100_nat_squared() int64 {
	var sum int64 = 0
	for i := 1; i <= 100; i++ {
		sum += int64(i)
	}

	return sum * sum
}
