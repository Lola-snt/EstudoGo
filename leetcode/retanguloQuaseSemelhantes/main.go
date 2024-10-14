package main

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func nearlySimilarRectangles(sides [][]int64) int64 {
	ratios := make(map[[2]int]int)

	for _, side := range sides {
		a, b := side[0], side[1]
		if a == 0 || b == 0 {
			continue // ignora lados inválidos
		}
		divisor := gcd(a, b)
		simplified := [2]int{int(a / divisor), int(b / divisor)}
		ratios[simplified]++
	}

	count := 0
	for _, occurrences := range ratios {
		if occurrences > 1 {
			count += occurrences * (occurrences - 1) / 2
		}
	}

	return int64(count)
}
