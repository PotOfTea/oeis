package output

import (
	"fmt"
	"oeis/pkg/consts"
)

func PrintError(msg string) {
	fmt.Println(string(consts.ColorRed), msg, string(consts.ColorReset))
}

func PrintResults(results []string, count int) {
	fmt.Printf("Found %v results. Showing first five:\n", count)
	for i := 0; i < len(results); i++ {
		fmt.Printf("%v) %v \n", i+1, results[i])
	}
}
