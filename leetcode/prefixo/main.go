//14. Prefixo comum mais longo
//
//Escreva uma função para encontrar a maior string de prefixo comum entre uma matriz de strings.
//
//Se não houver um prefixo comum, retorne uma string vazia

package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"a"}))
}

func longestCommonPrefix(strs []string) string {
	prencontrado := ""
	prparcial := ""

	if len(strs) == 0 {
		return ""
	}

	for tampr := 1; tampr <= len(strs[0]); tampr++ {
		if len(strs[0]) < tampr {
			return prencontrado
		}
		pr := strs[0][:tampr]
		for i := 0; i < len(strs); i++ {
			if len(strs[i]) < tampr {
				return prencontrado
			}
			if sub := strs[i][:tampr]; sub == pr {
				prparcial = sub
			} else {
				return prencontrado
			}
		}
		prencontrado = prparcial
	}
	return prencontrado
}
