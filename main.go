package main

import (
	"fmt"
	"markdown/content"
	"markdown/load"
	"sync"

	"github.com/gammazero/workerpool"
	"github.com/spf13/afero"
)

var directories = map[string]string{
	"data":      "./_source/data/",
	"images":    "./_source/images/",
	"converted": "./_source/converted/",
	"posts":     "./_source/posts/",
}

var poolSize = 20

func main() {
	var mutex = &sync.Mutex{}

	fmt.Printf("Starting markdown generation..\n\n")
	fs := afero.NewOsFs()
	out := load.Load(fs, directories["converted"]+"out.json")

	wp := workerpool.New(poolSize)

	// write markdown file stubs
	for _, item := range out.All {
		i := item
		wp.Submit(func() {
			file, mkd := content.Post(i, out, directories["posts"])
			mutex.Lock()
			content.Write(fs, file, mkd)
			mutex.Unlock()
		})
	}

	wp.StopWait()
	fmt.Printf("\nEnding markdown..\n")
}
