package day2

import (
	"log"
)

func Part1() int {
	depth := 0
	position := 0
	for _, v := range getDirections() {
		switch v.direction {
		case "forward":
			position += v.distance
		case "down":
			depth += v.distance
		case "up":
			depth -= v.distance
		default:
			log.Fatal("bad direction")
		}
	}
	return depth * position
}

func Part2() int {
	aim := 0
	depth := 0
	position := 0
	for _, v := range getDirections() {
		switch v.direction {
		case "forward":
			position += v.distance
			depth += aim * v.distance
		case "down":
			aim += v.distance
		case "up":
			aim -= v.distance
		}
	}
	return depth * position
}
