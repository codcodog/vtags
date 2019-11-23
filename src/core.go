package vtags

import (
	"fmt"
	"os"
)

func Run(opts *Options) {
	if len(opts.Source) == 0 {
		fmt.Fprintf(os.Stderr, "no files specified. try \"vtags --help\"\n")
		os.Exit(exitErr)
	}

	tags := parse(opts)
	fmt.Println(tags)
}
