package main

// exo 1

func reverseString(str string) string {
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	return reversed

}

// exo 2

func sumNumbers(n int) int {
	if n <= 0 {
		return 0
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum

}

// exo 3

func EvenOrOdd(n int) string {
	if n%2 == 0 {
		return "even"
	}
	return "odd"
}

//ou

func IsEven(n int) bool {
	return n%2 == 0
}

// exo 4

func CountLetters(str string) map[rune]int {
	count := make(map[rune]int)
	for _, char := range str {
		count[char]++
	}
	return count
}

// exo 5

func Fibonacci(n int) []int {
	fib := []int{0, 1}
	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	return fib[:n]
}

func main() {
	// exo 1
	println(reverseString("abcde"))

	// exo 2
	println(sumNumbers(5))

	// exo 3
	println(EvenOrOdd(4))
	println(IsEven(5))

	// exo 4
	count := CountLetters("hello")
	for char, cnt := range count {
		println(string(char), cnt)
	}

	// exo 5
	fib := Fibonacci(10)
	for _, num := range fib {
		print(num, " ")
	}
}
