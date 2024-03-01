package rotationalcipher

func RotationalCipher(plain string, shiftKey int) (cipher string) {
	for _, charRune := range plain {
		if 'a' <= charRune && charRune <= 'z' {
			cipher += string('a' + ((charRune + rune(shiftKey) - 'a') % 26))
		} else {
			cipher += string('A' + ((charRune + rune(shiftKey) - 'A') % 26))
		}
	}
	return

}
