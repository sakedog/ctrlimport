package ctrlimport

import (
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "ctrlimport finds imports are not allowed"

var (
	oks []string
	ngs []string
)

func init() {
	Analyzer.Flags.Func("ok", "allowed import paths", func(ok string) error { oks = append(oks, ok); return nil })
	Analyzer.Flags.Func("ng", "prohibited import paths", func(ng string) error { ngs = append(ngs, ng); return nil })
}

// Analyzer finds imports are not allowed
var Analyzer = &analysis.Analyzer{
	Name:     "ctrlimport",
	Doc:      doc,
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	ctrl := parse(oks, ngs)

	for _, f := range pass.Files {
		for _, ip := range f.Imports {
			path, err := strconv.Unquote(ip.Path.Value)
			if err != nil {
				return nil, err
			}

			if !ctrl.handle(path) {
				pass.Reportf(ip.Pos(), "import %s is not allowed", path)
			}
		}
	}
	return nil, nil
}

type ctrl struct {
	okFullPath      []string
	ngFullPath      []string
	okRecursionPath []string
	ngRecursionPath []string
}

func parse(oks, ngs []string) *ctrl {
	c := &ctrl{}

	c.okFullPath, c.okRecursionPath = parsePath(oks)
	c.ngFullPath, c.ngRecursionPath = parsePath(ngs)

	return c
}

func parsePath(paths []string) (full, recursion []string) {
	for _, path := range paths {
		if isRecusion(path) {
			recursion = append(recursion, parentPath(path))
			continue
		}

		full = append(full, path)
	}

	return full, recursion
}

func isRecusion(path string) bool {
	return filepath.Base(path) == "..."
}

func parentPath(path string) string {
	return filepath.Dir(path)
}

func (c *ctrl) handle(path string) bool {
	for _, okPath := range c.okFullPath {
		if okPath == path {
			return true
		}
	}

	for _, ngPath := range c.ngFullPath {
		if ngPath == path {
			return false
		}
	}

	for _, okPath := range c.okRecursionPath {
		if containPath(path, okPath) {
			return true
		}
	}

	for _, ngPath := range c.ngRecursionPath {
		if containPath(path, ngPath) {
			return false
		}
	}

	return true
}

func containPath(target, parent string) bool {
	return strings.HasPrefix(target, parent)
}
