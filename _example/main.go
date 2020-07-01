package main

import (
	"fmt"

	"github.com/tobiaszheller/strmod"
)

func main() {
	in := "Some text/Other text/Something else "

	out := strmod.Modify(in,
		strmod.SplitAndReturnLastPart("/"),
		strmod.TrimSpace(),
		strmod.ReplaceAll(" ", "_"),
		strmod.ToLower(),
	)
	fmt.Println(out)
}
