package printconsole

import (
	"fmt"

	. "github.com/logrusorgru/aurora"
)

func PrintHint(a ...interface{}) {
	fmt.Println(Gray(12, a))
}

func PrintWarning(a ...interface{}) {
	fmt.Println(Yellow(a))
}

func PrintError(a ...interface{}) {
	fmt.Println(Red(a))
}