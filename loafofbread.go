package piscine

func LoafOfBread(str string) string {
	if len(str) > 0 && len(str) < 5 {
		return "Invalid Output\n"
	}
	if len(str) == 0 {
		return "\n"
	}
	result := ""
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		}
		count++
		result += string(str[i])
		if count == 5 {
			result += " "
			count = 0
			i++
		}
	}
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	return result + "\n"
}
