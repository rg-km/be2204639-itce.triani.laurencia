import "fmt"

func main() {
	var students = []int{1, 1, 1, 0, 0, 1}
	var sandwiches = []int{1, 0, 0, 0, 1, 1}
	fmt.Println(CountStudents(students, sandwiches))
}

func CountStudents(students []int, sandwiches []int) int {
	var count int
	var counter [2]int
	for i := 0; i < len(students); i++ {
		counter[students[i]]++
	}
	for i := 0; i < len(sandwiches); i++ {
		if sandwiches[i] == 0 {
			counter[0]--
		} else {
			counter[1]--
		}
		if counter[0] < 0 || counter[1] < 0 {
			count++
		}
	}
	return count
	// TODO: replace this
}