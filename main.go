package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

//partie 1
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

//partie 2
// exo 1

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

// exo 2

type User struct {
	User string `json:"username"`
	Age  int    `json:"age"`
}

func (u User) ToJSON() string {
	b, _ := json.Marshal(u)
	return string(b)
}

// exo 3

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello, World!"})
}

// exo 4

func MapFilter(m map[string]int) map[string]int {
	result := make(map[string]int)
	for k, v := range m {
		if v%2 == 0 {
			result[k] = v
		}
	}
	return result
}

// exo 5

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {

	// pour print ceux de la partie 2. à masquer celui de la partie 1 pour éviter les conflits
	// exo 1
	p := Person{Name: "Danh"}
	println(p.Greet())

	// exo 2
	u := User{User: "danhaxe", Age: 25}
	println(u.ToJSON())

	// exo 3
	http.HandleFunc("/hello", HelloHandler)
	http.ListenAndServe(":8080", nil)

	// exo 4
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	filtered := MapFilter(m)
	for k, v := range filtered {
		println(k, v)
	}
	// exo 5
	result, err := divide(10, 2)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Result:", result)
	}
}
