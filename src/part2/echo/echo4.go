package echo

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "add new line to the end")
var sep = flag.String("s", " ", "separate by character")

func Echo4() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))

	if *n {
		fmt.Println()
	}
}
