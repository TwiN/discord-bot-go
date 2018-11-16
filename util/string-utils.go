package util


func PadRight(s string, expectedLength int, pad string) string {
	for {
		s += pad
		if len(s) > expectedLength {
			return s[0:expectedLength]
		}
	}
}
