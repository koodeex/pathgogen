package core

func IsStringInTheArray(target string, array []string) bool {
	for _, str := range array {
		if target == str {
			return true
		}
	}
	return false
}
