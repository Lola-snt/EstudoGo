package main

type StringWithRemainder struct {
	Value     string
	Length    int
	Remainder int
}

func RemainderSorting(strArr []string) []string {
	var stringsWithRemainders []StringWithRemainder

	for _, s := range strArr {
		length := len(s)
		remainder := length % 3
		stringsWithRemainders = append(stringsWithRemainders, StringWithRemainder{s, length, remainder})
	}

	for i := 1; i < len(stringsWithRemainders); i++ {
		key := stringsWithRemainders[i]
		j := i - 1

		for j >= 0 && (stringsWithRemainders[j].Remainder > key.Remainder ||
			(stringsWithRemainders[j].Remainder == key.Remainder && stringsWithRemainders[j].Value > key.Value)) {
			stringsWithRemainders[j+1] = stringsWithRemainders[j]
			j--
		}
		stringsWithRemainders[j+1] = key
	}

	var sortedStrings []string
	for _, str := range stringsWithRemainders {
		sortedStrings = append(sortedStrings, str.Value)
	}
	return sortedStrings
}
