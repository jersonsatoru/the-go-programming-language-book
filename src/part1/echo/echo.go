package echo

import (
	"fmt"
	"os"
	"strings"
)

func Echo() {
	fmt.Println(strings.Join(os.Args[1:], ","))
}
