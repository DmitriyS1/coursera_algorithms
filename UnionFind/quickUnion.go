package main

// With this function we have logN time complexity
func root(i int, unionArr []int) int {
	for i != unionArr[i] {
		i = unionArr[i]
	}

	return i
}

// With this function we have log * N complexity
func rootWithCompression(i int, unionArr []int) int {
	for i != unionArr[i] {
		unionArr[i] = unionArr[unionArr[i]]
		i = unionArr[i]
	}

	return i
}

func Connected(p int, q int, unionArr []int) bool {
	return rootWithCompression(p, unionArr) == root(q, unionArr)
}

func Unite(p int, q int, unionArr *[]int) {
	i := rootWithCompression(p, *unionArr)
	j := rootWithCompression(q, *unionArr)

	(*unionArr)[i] = j
}

func WheightedUnite(p int, q int, unionArr *[]int, szArr *[]int) {
	i := rootWithCompression(p, *unionArr)
	j := rootWithCompression(q, *unionArr)

	if i == j {
		return
	}

	if (*szArr)[i] < (*szArr)[j] {
		(*unionArr)[i] = j
		(*szArr)[j] += (*szArr)[i]
	} else {
		(*unionArr)[j] = i
		(*szArr)[i] += (*szArr)[j]
	}
}