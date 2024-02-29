package phoneNumber

func getNumbers(input string) string {
	var number string

	for _, t := range input {
		if t >= 48 && t <= 57 {
			number += string(t)
		}
	}

	return number
}
