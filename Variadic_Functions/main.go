package main

// variadic function is a function that can take a variable number of arguments of the same type. In Go, you can define a variadic function by using an ellipsis (...) before the type of the last parameter. This allows you to pass any number of arguments of that type to the function.
func add(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func greet(prefix string, names ...string) {
	for _, name := range names {
		println(prefix, name)
	}
}

func main() {
	sum := add(10, 20, 30)
	println(sum)

	mps :=[]string{"Jon", "Arya", "Sansa"}
	greet("Hello, ", mps...)
}