package decoration

import (
	"fmt"
	"strings"
)

func SetTheme(str string) {
	fmt.Println("<<-- ", str, " -->>")
	fmt.Println("<<-------------------------------->>")
}

func SetSubTheme(str string) {
	fmt.Println()
	fmt.Println("*", strings.ToUpper(str))
	fmt.Println("********")
}
