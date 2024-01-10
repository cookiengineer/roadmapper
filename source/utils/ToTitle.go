package utils

func ToTitle(value string) string {

	var result string

	for v := 0; v < len(value); v++ {

		character := string(value[v])

		if character >= "0" && character <= "9" {
			result += character
		} else if character >= "A" && character <= "Z" {
			result += character
		} else if character >= "a" && character <= "z" {
			result += character
		} else if character == "(" {
			result += character
		} else if character == ")" {
			result += character
		} else if character == ":" {
			result += character
		} else if character == "." {
			result += character
		} else if character == "_" {
			result += character
		} else if character == "-" {
			result += character
		} else if character == "+" {
			result += character
		} else if character == " " {
			result += character
		} else {
			continue
		}

	}

	return result

}
