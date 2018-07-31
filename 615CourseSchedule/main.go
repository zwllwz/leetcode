package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make([]int, numCourses)
	graph := map[int][]int{}
	for _, prereq := range prerequisites {
		a := prereq[1]
		b := prereq[0]
		// 1 -> 0, [0, 1]
		inDegree[b]++
		graph[a] = append(graph[a], b)
	}

	queue := []int{}

	for i := range inDegree {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		course := queue[0]
		queue = queue[1:]
		for _, n := range graph[course] {
			inDegree[n] = inDegree[n] - 1
			if inDegree[n] == 0 {
				queue = append(queue, n)
			}
		}
	}

	for i := range inDegree {
		if inDegree[i] != 0 {
			return false
		}
	}

	return true
}

func main() {

}
