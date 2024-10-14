package main

func bubbleSort(arr []string) []string {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				// Troca
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func mostActive(customers []string) []string {
	tradeCount := make(map[string]int)

	// Contar o número de trades por cliente
	for _, customer := range customers {
		tradeCount[customer]++
	}

	totalTrades := len(customers)
	threshold := totalTrades * 5 / 100 // 5% do total

	activeCustomers := []string{}

	// Filtrar clientes que atendem ao critério
	for name, count := range tradeCount {
		if count >= threshold {
			activeCustomers = append(activeCustomers, name)
		}
	}

	// Ordenar a lista alfabeticamente usando bubble sort
	activeCustomers = bubbleSort(activeCustomers)

	return activeCustomers
}
