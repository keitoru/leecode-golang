package main

import "fmt"

/*func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make(map[int]int, numCourses)
	m := make(map[int][]int)

	for i := 0; i < len(prerequisites); i++ {
		inDegree[prerequisites[i][0]]++
		if _, ok := m[prerequisites[i][1]]; !ok {
			m[prerequisites[i][1]] = []int{}
		}
		m[prerequisites[i][1]] = append(m[prerequisites[i][1]], prerequisites[i][0])
	}

	var queue []int
	for i := 0; i < len(inDegree); i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	count := 0
	for len(queue) > 0 {
		selected := queue[0]
		queue = queue[1:]

		count++
		toEnQueue := m[selected]
		if len(toEnQueue) > 0 {
			for i := 0; i < len(toEnQueue); i++ {
				inDegree[toEnQueue[i]]--
				if inDegree[toEnQueue[i]] == 0 {
					queue = append(queue, toEnQueue[i])
				}
			}
		}
	}

	return count == numCourses
}*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	indegrees := make(map[int]int, numCourses)

	m := make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		m[i] = []int{}
	}

	for _, cp := range prerequisites {
		// 入度个数
		indegrees[cp[0]]++
		// 出度
		m[cp[1]] = append(m[cp[1]], cp[0])
	}

	var queue []int
	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		pre := queue[0]
		queue = queue[1:]
		numCourses--

		for _, cur := range m[pre] {
			indegrees[cur]--
			if indegrees[cur] == 0 {
				queue = append(queue, cur)
			}
		}
	}

	return numCourses == 0
}

func main() {
	p := [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}
	res := canFinish(5, p)
	fmt.Println(res)
}
