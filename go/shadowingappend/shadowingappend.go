package shadowingappend

// const make = 1
// const append = 2

func append([]int, int) []int {
	return []int{1, 1, 1}
}

func main() {
	array := make([]int, 0, 10)
	array = append(array, 1)
}
