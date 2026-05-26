package main

import "fmt"

func CheckAllZero(store map[rune]int) bool {
	for _, val := range store {
		if val != 0 {
			return false
		}
	}
	return true
}

func CountAnagramsOccurence(txt string, pat string) int {
	result := 0
	store := make(map[rune]int)
	patArray := []rune(pat)
	txtArray := []rune(txt)
	for _, ch := range patArray {
		store[ch]++
	}

	i := 0
	j := 0
	n := len(txtArray)
	size := len(patArray)

	for j < n {
		_, ok := store[txtArray[j]]
		if ok {
			store[txtArray[j]]--
		}

		if j-i+1 == size {
			if CheckAllZero(store) {
				result++
			}
			_, ok := store[txtArray[i]]
			if ok {
				store[txtArray[i]]++
			}
			i++
		}
		j++
	}
	return result
}

func main() {
	txt := "forxxorfxdofr"
	pat := "for"
	count := CountAnagramsOccurence(txt, pat)
	fmt.Println("Count is: ", count)

}
