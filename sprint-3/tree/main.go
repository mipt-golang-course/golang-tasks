package main

/*
Course `Web services on Go`, week 1, homework, `tree` program.
See: week_01\materials.zip\week_1\99_hw\tree

mkdir -p week01_homework/tree
pushd week01_homework/tree
go mod init tree
go mod tidy
pushd ..
go work init
go work use ./tree/
go vet tree
gofmt -w tree
go test -v tree
go run tree . -f
go run tree ./tree/testdata
cd tree && docker build -t mailgo_hw1 .

https://en.wikipedia.org/wiki/Tree_(command)
https://mama.indstate.edu/users/ice/tree/
https://stackoverflow.com/questions/32151776/visualize-tree-in-bash-like-the-output-of-unix-tree

*/

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

/*
	Example output:

	├───project
	│	└───gopher.png (70372b)
	├───static
	│	├───a_lorem
	│	│	├───dolor.txt (empty)
	│	├───css
	│	│	└───body.css (28b)
	...
	│			└───gopher.png (70372b)

	- path should point to a directory,
	- output all dir items in sorted order, w/o distinction file/dir
	- last element prefix is `└───`
	- other elements prefix is `├───`
	- nested elements aligned with one tab `	` for each level
*/

const (
	EOL             = "\n"
	BRANCHING_TRUNK = "├───"
	LAST_BRANCH     = "└───"
	TRUNC_TAB       = "│\t"
	LAST_TAB        = "\t"
	EMPTY_FILE      = "empty"
	ROOT_PREFIX     = ""

	USE_RECURSION_ENV_KEY = "RECURSIVE_TREE"
	USE_RECURSION_ENV_VAL = "YES"
)

type treeWorker struct {
	printFiles bool
	out        io.Writer
}

func NewTreeWorker(out io.Writer, printFiles bool) *treeWorker {
	return &treeWorker{
		printFiles: printFiles,
		out:        out,
	}
}

func (tw *treeWorker) prepareDirs(path string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	processedEntries := make([]os.DirEntry, 0, len(entries))
	for i := 0; i < len(entries); i++ {
		if shouldInclude(entries[i], tw.printFiles) {
			processedEntries = append(processedEntries, entries[i])
		}
	}

	sortDirsByName(processedEntries)
	return processedEntries, nil
}

func shouldInclude(entry os.DirEntry, printFiles bool) bool {
	if !printFiles && !entry.IsDir() {
		return false
	}
	return entry.Name() != ".DS_Store"
}

func sortDirsByName(entries []os.DirEntry) {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})
}

type dirWorker struct {
	entry      os.DirEntry
	worker     *treeWorker
	basePrefix string
	branch     string
	nextPrefix string
}

func NewDirWorker(worker *treeWorker, entry os.DirEntry, isLast bool, basePrefix string) *dirWorker {
	dw := &dirWorker{
		entry:      entry,
		worker:     worker,
		basePrefix: basePrefix,
	}

	if isLast {
		dw.branch = LAST_BRANCH
		dw.nextPrefix = basePrefix + LAST_TAB
	} else {
		dw.branch = BRANCHING_TRUNK
		dw.nextPrefix = basePrefix + TRUNC_TAB
	}

	return dw
}

func (dw *dirWorker) getFileSizeString() (string, error) {
	if dw.entry.IsDir() {
		return "", nil
	}

	info, err := dw.entry.Info()
	if err != nil {
		return "", err
	}

	if info.Size() == 0 {
		return " (" + EMPTY_FILE + ")", nil
	}

	return fmt.Sprintf(" (%db)", info.Size()), nil
}

func (dw *dirWorker) printLine() error {
	sizeStr, err := dw.getFileSizeString()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(dw.worker.out, strings.Join([]string{dw.basePrefix, dw.branch, dw.entry.Name(), sizeStr}, ""))
	return err
}

func (tw *treeWorker) walkOnTree(path, prefix string) error {
	entries, err := tw.prepareDirs(path)
	if err != nil {
		return err
	}

	for i, entry := range entries {
		processor := NewDirWorker(tw, entry, i == len(entries)-1, prefix)

		if err := processor.printLine(); err != nil {
			return err
		}

		if processor.entry.IsDir() {
			err := tw.walkOnTree(filepath.Join(path, processor.entry.Name()), processor.nextPrefix)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	// This code is given
	if len(os.Args) != 2 && len(os.Args) != 3 {
		panic("usage: go run main.go . [-f]")
	}

	out := os.Stdout
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

// dirTree: `tree` program implementation, top-level function, signature is fixed.
// Write `path` dir listing to `out`. If `printFiles` is set, files is listed along with directories.
func dirTree(out io.Writer, path string, printFiles bool) error {
	// Function to implement, signature is given, don't touch it.

	worker := NewTreeWorker(out, printFiles)
	return worker.walkOnTree(path, ROOT_PREFIX)
}
