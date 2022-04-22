# golang_valid_palindrone

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

## Examples

```
Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
```

## 題目分析

已知需要判斷所有字串 alphanumberic 字元是否滿足迴文

初步想法

先把 alphanumeric 挑出來成為一個字串 alphaNum

然後再 判斷每個對應的值, i, (len(alphaNum)-1) - i 是否相等


## 參考別人的作法

可以從字串兩邊同時找尋

假設 lp 代表左邊 index, rp 代表右邊 index

每次先判斷 lp , rp 是否為 alphanumeric

如果 lp 不是 則把 lp 向右移動 1 格 重新在比一次

如果 rp 不是 則把 rp 向左移動 1 格 重新在比一次

## 實作

```golang
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

```