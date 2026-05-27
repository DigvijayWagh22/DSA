package main

func CheckAllZero(store map[rune]int) bool {
	for _, val := range store {
		if val != 0 {
			return false
		}
	}
	return true
}

func FindAllAnagrams(txt string, pat string) []int {
	result := []int{}
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
				result = append(result, i)
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

// txt := "cbaebabacd"
// pat := "abc"
// allAnagrams := FindAllAnagrams(txt, pat)
// for _, val := range allAnagrams {
// 	fmt.Println(val)
// }
