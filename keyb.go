package keyb

import (
	"fmt"
)

func WaitKeyEnter() {
	var input string
	fmt.Scanln(&input)
}
