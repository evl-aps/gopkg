package decoration

import (
	"fmt"
	"strings"
)

// colorRed := "\033[31m"
// colorYellow := "\033[33m"

func SetTheme(str string) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"

	fmt.Println(string(colorGreen), "<<-- ", str, " -->>")
	fmt.Println("<<-------------------------------->>", string(colorReset))
}

func SetSubTheme(str string) {
	colorReset := "\033[0m"
	colorCyan := "\033[36m"

	fmt.Println()
	fmt.Println(string(colorCyan), "*", strings.ToUpper(str))
	fmt.Println("********", string(colorReset))
}
