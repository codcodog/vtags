package vtags

import (
	"fmt"
	"os"
	"strings"
)

// exit code
const (
	exitOk  = 0
	exitErr = 2
)

const usage = `Usage: vtags [options] [source_file(s)]

OPTIONS
	--exclude=[pattern]	Add pattern to a list of excluded files and directories.

	-R recurse into directories in the file list.

	-es5 xxx

	+es5 xxx

	-es6 xxx

	+es6 xxx

`

type Options struct {
	Exclude []string
	Recurse bool
	Es5     bool
	Es6     bool
	Source  []string
}

func defaultOptions() *Options {
	return &Options{
		Es5: true,
	}
}

func ParseOptions() *Options {
	opts := defaultOptions()
	parseOptions(opts, os.Args[1:])

	return opts
}

func parseOptions(opts *Options, args []string) {
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "-h", "--help":
			help(exitOk)
		case "-r", "-R":
			opts.Recurse = true
		case "-es5":
			opts.Es5 = true
		case "+es5":
			opts.Es5 = false
		case "-es6":
			opts.Es6 = true
		case "+es6":
			opts.Es6 = false
		default:
			if match, value := optString(arg, "-e", "--exclude="); match {
				opts.Exclude = append(opts.Exclude, value)
				continue
			}

			prepareSource(opts, arg)
		}
	}
}

func help(code int) {
	os.Stdout.WriteString(usage)
	os.Exit(code)
}

func prepareSource(opts *Options, file string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "cannot open input file %q : no such file or directory\n", file)
		os.Exit(exitErr)
	}
	opts.Source = append(opts.Source, file)
}

func optString(arg string, prefixes ...string) (bool, string) {
	for _, prefix := range prefixes {
		if strings.HasPrefix(arg, prefix) {
			return true, arg[len(prefix):]
		}
	}

	return false, ""
}
