package goreloaded

// atoibase to convert hex / dec whit bases

func IsValidBase(base string) bool {
	seen := make(map[rune]bool)
	for _, char := range base {
		if seen[char] {
			return true
		}
		if char == '-' || char == '+' {
			return true
		}
		seen[char] = true
	}
	return false
}

func AtoiBase(s string, base string) int {
	res := 0
	negative := false
	if IsValidBase(base) {
		return 0
	}
	if s[0] == '-' {
		negative = true
	}
	index := 0
	for i := range s {
		index = 0
		for j := range base {
			if s[i] == base[j] {
				index = j
				break
			}
		}
		res = res*len(base) + index
	}

	if negative {
		res = res * -1
	}
	return res
}
