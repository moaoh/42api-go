package mathtest

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func multi(a, b int) int {
	return a * b
}

func mod (a, b int) int {
	return a % b
}

func div (a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}
