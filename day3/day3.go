package day3

import "fmt"

func Part1() int {
	gamma := 0
	epsilon := 0
	m := map[int]int{}
	for _, v := range input {
		for k, v := range v {
			if v == '1' {
				m[k] += 1
			}
		}
	}
	for i := 0; i < len(m); i++ {
		if m[i] > len(input)/2 {
			gamma = gamma | (1 << (len(m) - (i + 1)))
		} else {
			epsilon = epsilon | (1 << (len(m) - (i + 1)))
		}
	}
	return gamma * epsilon
}

func Part2() {
	ox := 0
	co := 0

	m := genMap()
	mm := genMap()

	max := ""
	min := ""
	for i := 0; i < len(input[0]); i++ {
		one, zero := 0, 0
		for k := range m {
			if k[i] == '1' {
				one++
			} else {
				zero++
			}
		}
		if one >= zero {
			max += "1"
		} else {
			max += "0"
		}
		m = filter(max, m)
		if len(m) == 1 {
			break
		}
	}
	for i := 0; i < len(input[0]); i++ {
		one, zero := 0, 0
		for k := range mm {
			if k[i] == '1' {
				one++
			} else {
				zero++
			}
		}
		if one < zero {
			min += "1"
		} else {
			min += "0"
		}
		mm = filter(min, mm)
		if len(mm) == 1 {
			break
		}
	}
	for _, v := range m {
		ox = v
	}
	for _, v := range mm {
		co = v
	}
	fmt.Println(ox * co)
}

func filter(s string, m map[string]int) map[string]int {
	mask := maskInt(s)
	out := map[string]int{}
	for k, v := range m {
		if (v>>(len(k)-len(s)) ^ mask) == 0 {
			out[k] = v
		}

		//could have done this but this is boring :)
		// if s == k[:len(s)-1] {
		// 	out[k] = v
		// }
	}
	return out
}

func maskInt(s string) int {
	out := 0
	for k, v := range s {
		if v == '1' {
			out = out | (1 << (len(s) - (k + 1)))
		}
	}
	return out
}

func genMap() map[string]int {
	m := map[string]int{}
	for _, v := range input {
		m[v] = maskInt(v)
	}
	return m
}
