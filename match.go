package bracket

func Check(input string) bool {
	if input == "{" || input == "}" {
		return false
	}
	return true
}
