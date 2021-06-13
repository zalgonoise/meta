package main

import (
	"fmt"

	cm "base/strings/charmatcher"
)

func main() {
	matcher, err := cm.NewCharMatcher(`abcdefg`)

	if err != nil {
		panic(err)
	}

	if matcher.Contains(`abc`) {
		fmt.Println(`yep!`)
	} else {
		fmt.Println(`nope!`)
	}
}