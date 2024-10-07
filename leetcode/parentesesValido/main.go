package main

import "fmt"

func main() {
	fmt.Println(isValid("){"))
}

func isValid(s string) bool {
	parentesesAberto := 0
	colchetesAberto := 0
	chavesAberta := 0
	if len(s) <= 1 {
		return false
	}
	for i := 0; i <= len(s)-2; i += 1 {
		atual := s[i]
		proxima := s[i+1]
		switch atual {
		case '(':
			parentesesAberto++
			if proxima == '}' || proxima == ']' {
				return false
			}
			break
		case '[':
			colchetesAberto++
			if proxima == '}' || proxima == ')' {
				return false
			}
			break
		case '{':
			chavesAberta++
			if proxima == ']' || proxima == ')' {
				return false
			}
			break
		case ')':
			if parentesesAberto < 0 {
				return false
			}
			break
		case '}':
			if chavesAberta < 0 {
				return false
			}
			break
		case ']':
			if colchetesAberto < 0 {
				return false
			}
			break
		}
		switch proxima {
		case ')':
			parentesesAberto--
			if parentesesAberto < 0 {
				return false
			}
			break
		case '}':
			chavesAberta--
			if chavesAberta < 0 {
				return false
			}
			break
		case ']':
			colchetesAberto--
			if colchetesAberto < 0 {
				return false
			}
			break

		}
	}
	return parentesesAberto == 0 && colchetesAberto == 0 && chavesAberta == 0

}
