package vtags

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var tags bytes.Buffer

func parse(opts *Options) string {
	files := getFilesName(opts.Source, opts.Exclude)

	for _, file := range files {
		f, err := os.Open(file)
		checkErr(err)

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			generateTags(file, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			os.Exit(exitErr)
		}

		f.Close()
	}

	return tags.String()
}

func getFilesName(paths []string, exclude []string) []string {
	var files []string

	for _, path := range paths {
		filesTmp := getFileList(path, exclude)
		files = append(files, filesTmp...)
	}

	return files
}

func getFileList(path string, exclude []string) []string {
	var files []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		checkErr(err)

		if len(exclude) > 0 && matchExcludRegexp(path, exclude) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if isVue(path) {
			files = append(files, path)
		}

		return nil
	})
	checkErr(err)

	return files
}

func matchExcludRegexp(path string, exclude []string) bool {
	for _, pattern := range exclude {
		pat := "^" + pattern
		re := regexp.MustCompile(pat)

		if re.MatchString(path) {
			return true
		}
	}

	return false
}

func isVue(file string) bool {
	return strings.HasSuffix(file, ".vue")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(exitErr)
	}
}

func debug(v ...interface{}) {
	for _, value := range v {
		fmt.Printf("%+v \n", value)
	}
	os.Exit(1)
}

func generateTags(file, line string) {
	exportParse(file, line)
	propsParse(file, line)
	dataParse(file, line)
	functionParse(file, line)
	createdParse(file, line)
	computedParse(file, line)
}

func exportParse(file, line string) {
	pattern := `^[\s]*export[\s]*default[\s]*{$`
	matchRegForTag(pattern, file, line, "export", "v")
}

func propsParse(file, line string) {
	pattern := `^[\s]*props:*`
	matchRegForTag(pattern, file, line, "props", "v")
}

func dataParse(file, line string) {
	pattern := `^[\s]*data[\s]*([\s]*)[\s]*{$`
	matchRegForTag(pattern, file, line, "data", "v")
}

func functionParse(file, line string) {
	pattern := `^[\s]*[a-zA-Z0-9_]+[\s]*\(*\)[\s]*{$`
	matchRegForTag(pattern, file, line, "function", "f")
}

func createdParse(file, line string) {
	pattern := `^[\s]*created[\s]*\([\s]*\)[\s]*{$`
	matchRegForTag(pattern, file, line, "created", "v")
}

func computedParse(file, line string) {
	pattern := `^[\s]*computed:[\s]*{$`
	matchRegForTag(pattern, file, line, "computed", "v")
}

func matchRegForTag(pattern, file, line, keyword, define string) {
	re := regexp.MustCompile(pattern)
	if re.MatchString(line) {
		matchPattern := "/^" + line + "$" + `/;"`
		tag := fmt.Sprintf("%s\t%s\t%s\t%s\n", keyword, file, matchPattern, define)
		tags.WriteString(tag)
	}
}
