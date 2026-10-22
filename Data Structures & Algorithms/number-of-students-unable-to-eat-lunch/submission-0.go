func countStudents(students []int, sandwiches []int) int {
	var consMisses int
    for len(students) > consMisses {
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			consMisses = 0
		} else {
			students = append(students[1:], students[0])
			consMisses++
		}
	}
	return len(students)
}