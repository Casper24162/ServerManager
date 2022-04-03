package format

func Format(colorName, text string) string {
	var color string
	switch {
	case colorName == "red":
		color = "\033[31m"
	case colorName == "green":
		color = "\033[32m"
	case colorName == "yellow":
		color = "\033[33m"
	case colorName == "blue":
		color = "\033[34m"
	case colorName == "purple":
		color = "\033[35m"
	case colorName == "cyan":
		color = "\033[36m"
	case colorName == "white":
		color = "\033[37m"
	default:
		color = "\033[33m"
	}
	return color + text + "\033[0m"
}
