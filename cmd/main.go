package main

import (
	"fmt"

	"github.com/caitouyun/rewriter"
)

func main() {
	rewriter := rewriter.NewRewriter([]rewriter.Rule{
		{"/old", "/new"},
	})

	fmt.Println(rewriter.MustRewrite("/api"))
}
