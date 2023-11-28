package main

import (
	blog "evo-blog-gf/internal/evo-blog-gf"
	_ "go.uber.org/automaxprocs"
	"os"
)

func main() {
	command := blog.GetBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
