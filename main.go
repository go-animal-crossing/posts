package main

import (
	"fmt"
)

var directories = map[string]string{
	"data":      "./_source/data/",
	"images":    "./_source/images/",
	"converted": "./_source/converted/",
	"markdown":  "./_source/markdown/",
}

func main() {
	fmt.Printf("Starting markdown generation..\n\n")
	//fs := afero.NewOsFs()

	fmt.Printf("\nEnding conversion..\n")
}
