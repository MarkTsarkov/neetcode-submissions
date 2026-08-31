func isAnagram(s string, t string) bool {
	if len(s)!=len(t) {
		return false
	}
	hashMap := make(map[rune]int)
	for _, l := range s {
		hashMap[l]++
	}
	for _, l := range t {
		if _, ok := hashMap[l]; ok {
			hashMap[l]--
		} else {
			return false
		}
	}
	for _, i := range hashMap{
		if i != 0 {
			return false
		}
	}
	return true
}
