package main

import "fmt"

func main() {
	maxItog, minItog := findMax(1, 2321, -123, 434, 1123, 0, 5452, 3333, 9999, 1123)
	fmt.Println("Максимальное знание в массиве: ", maxItog)
	fmt.Println("Минимальное знание в массиве: ", minItog)
}

func findMax(massive ...int) (int, int) {
	if len(massive) == 0 {
		return 0, 0
	}

	max := massive[0]
	for _, v := range massive {
		if max < v {
			max = v
		}
	}

	min := massive[0]
	for _, v := range massive {
		if min > v {
			min = v
		}
	}
	return max, min

}
