func isPalindrome(s string) bool {
	l := 0
	r := len(s)-1
	for {
		if l >= r {
			break
		}

		ch1 := checkAndLower(s[l])
		if ch1 == ""{
			l++
			continue
		}

		ch2 := checkAndLower(s[r])
		if ch2==""{
			r--
			continue
		}

		if ch1 != ch2 {
			return false
		}
		l++
		r--
	}
	return true
}

func checkAndLower(s byte) string {
	if ('a' <= s && s <= 'z') || ('A' <= s && s <='Z') || ('0' <= s && s<='9') {
			ch := s | 32
			return string(ch)
	}
	return ""
}