package output

import (
	"fmt"
	"oeis/pkg/consts"
)

func PrintError(msg ...interface{}) {
	fmt.Println(string(consts.ColorRed), msg, string(consts.ColorReset))
}

func PrintInfo(msg ...interface{}) {
	fmt.Println(string(consts.ColorInfo), msg, string(consts.ColorReset))
}
