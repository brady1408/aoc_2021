package day1

func Part1() int {
	increase := 0
	for k, v := range input {
		if k == 0 {
			continue
		}
		if input[k-1] < v {
			increase++
		}
	}
	return increase
}

func Part2() int {
	increase := 0
	for k := range input {
		if k < 3 {
			continue
		}
		if getPrev(k) < getCurrent(k) {
			increase++
		}
	}
	return increase
}

func getPrev(i int) int {
	return input[i-3] + input[i-2] + input[i-1]
}

func getCurrent(i int) int {
	return input[i-2] + input[i-1] + input[i]
}
