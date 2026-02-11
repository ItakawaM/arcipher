package ciphers

func reverseString(s string) string {
	reversed := []rune(s)
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return string(reversed)
}

func RailFenceEncryptMessage(message string, key int) string {
	if key == 1 {
		return message
	}

	if key >= len(message) {
		return reverseString(message)
	}

	cycle := 2 * (key - 1)
	encrypted := make([]rune, len(message))
	runes := []rune(message)

	index := 0
	for level := key - 1; level >= 0; level-- {
		for charIndex := level; charIndex < len(runes); charIndex += cycle {
			encrypted[index] = runes[charIndex]
			index += 1

			// if middle row
			secondCharIndex := charIndex + cycle - 2*level
			if level != key-1 && level != 0 && secondCharIndex < len(message) {
				encrypted[index] = runes[secondCharIndex]
				index += 1
			}
		}
	}

	return string(encrypted)
}

func RailFenceDecryptMessage(encrypted string, key int) string {
	panic("Not Implemented!")
}
