func isPalindrome(s string) bool {
	pal := make([]byte, 0)
	for i := range s {
		if ('a' <= s[i] && s[i] <= 'z') || ('A' <= s[i] && s[i] <='Z') || ('0' <= s[i] && s[i]<='9') {
			ch := s[i] | 32
			pal=append(pal, ch)
		}
	}

	for i := range pal {
		if i == (len(pal))/2 {
			break
		}
		if pal[i] != pal[len(pal)-1-i] {
			return false
		}
	}
	return true
}
