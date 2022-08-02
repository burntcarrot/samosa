package samosa

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/cover"
)

var skipList map[string]struct{}

func populateSkipList() {
	skipList = map[string]struct{}{
		".git":     {},
		"testdata": {},
	}

}

func walkModDir(modfile string) ([]string, error) {
	d, _ := filepath.Split(modfile)
	file, err := os.Open(d)
	if err != nil {
		return nil, err
	}

	// read all dirs
	return file.Readdirnames(-1)

}

func getFileNames() ([]string, error) {
	fileDirs, err := walkDir()
	if err != nil {
		return nil, err
	}
	return fileDirs, nil
}

type Function struct {
	name      string
	startLine int
	startCol  int
	endLine   int
	endCol    int
}

type Visitor struct {
	fset    *token.FileSet
	name    string
	astFile *ast.File
	funcs   []*Function
}

// coverage returns the number of covered and total lines.
func (f *Function) coverage(profile *cover.Profile) (int, int) {
	total := 0
	covered := 0

	for _, b := range profile.Blocks {
		if b.StartLine > f.endLine || (b.StartLine == f.endLine && b.StartCol >= f.endCol) {
			break
		}

		if b.EndLine < f.startLine || (b.EndLine == f.startLine && b.EndCol <= f.startCol) {
			continue
		}

		total += b.NumStmt

		if b.Count > 0 {
			covered += b.NumStmt
		}
	}

	if total == 0 {
		total = 1
	}

	return covered, total
}

func getProfiles(coverageFilePath string) ([]*cover.Profile, error) {
	fmt.Print("starting to get profiles .....")
	profiles, err := cover.ParseProfiles(coverageFilePath)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

// getFunctions returns functions for a given file.
func getFunctions(filename string) ([]*Function, error) {
	fset := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		return nil, err
	}

	visitor := &Visitor{
		fset:    fset,
		name:    filename,
		astFile: parsedFile,
	}

	// traverse AST
	ast.Walk(visitor, visitor.astFile)

	return visitor.funcs, nil
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.FuncDecl:
		start := v.fset.Position(n.Pos())
		end := v.fset.Position(n.End())

		f := &Function{
			name:      n.Name.Name,
			startLine: start.Line,
			startCol:  start.Column,
			endLine:   end.Line,
			endCol:    end.Column,
		}

		v.funcs = append(v.funcs, f)
	}

	return v
}

type funcInfo struct {
	fileName       string
	pkgFileName    string
	functionName   string
	startLine      int
	endLine        int
	uncoveredLines int
}

func getFunctionInfo(profiles []*cover.Profile) ([]*funcInfo, int, int, error) {
	total := 0
	covered := 0
	var funcInfos []*funcInfo

	for _, profile := range profiles {
		filenames, err := getFileNames()
		if err != nil {
			return nil, 0, 0, err
		}
		fmt.Printf("acquired list of files for coverage report....")
		for _, filename := range filenames {
			functions, err := getFunctions(filename)
			if err != nil {
				return nil, 0, 0, err
			}
			for _, f := range functions {
				c, t := f.coverage(profile)

				fi := &funcInfo{
					fileName:       filename,
					pkgFileName:    filename,
					functionName:   f.name,
					startLine:      f.startLine,
					endLine:        f.endLine,
					uncoveredLines: t - c,
				}

				funcInfos = append(funcInfos, fi)
				total += t
				covered += c
			}

		}
	}

	return funcInfos, covered, total, nil
}

// walkDir gets list of go files in repo
func walkDir() ([]string, error) {
	var filesDir []string
	// get mod dir
	modeFile, err := getMod()
	if err != nil {
		return nil, err
	}
	dir, _ := filepath.Split(modeFile)
	log.Default().Println("mod file location:", dir)
	// walk dir
	if err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if skipList == nil {
			populateSkipList()
		}
		// filter files to go types
		p := filterFiles(path)
		if p != "" {
			filesDir = append(filesDir, p)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return filesDir, nil
}

func filterFiles(path string) string {
	extn := filepath.Ext(path)
	if extn == ".go" {
		if !strings.Contains(path, "test") {
			return path
		}
	}
	return ""
}
