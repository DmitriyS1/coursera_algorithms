package main

// Union Find is too slow because it has N^2 time complexity

func Connected(p int, q int, id_array []int) bool {
	return id_array[p] == id_array[q]
}

func Unite(p int, q int, id_array *[]int) {
	pid := (*id_array)[p]
	qid := (*id_array)[q]

	for i := 0; i < len(*id_array); i++ {
		if (*id_array)[i] == pid {
			(*id_array)[i] = qid
		}
	}
}
