package valid_palindrone

func isPalindrome(s string) bool {
	lp := 0
	rp := len(s) - 1
	for lp <= rp {
		// check isAlphaNum
		if !isAlphaNum(s[lp]) {
			lp++
			continue
		}
		if !isAlphaNum(s[rp]) {
			rp--
			continue
		}
		if toLower(s[rp]) != toLower(s[lp]) {
			return false
		}
		lp++
		rp--
	}
	return true
}

func isAlphaNum(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		b += ('a' - 'A')
	}
	return b
}
