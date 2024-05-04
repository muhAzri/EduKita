package utils

func ToSnakeCase(s string) string {
	var result string
	for i, c := range s {
		if 'A' <= c && c <= 'Z' {
			if i > 0 {
				result += "_"
			}
			result += string(c - 'A' + 'a')
		} else {
			result += string(c)
		}
	}
	return result
}
