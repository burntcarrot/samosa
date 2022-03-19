package internal

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/thoas/go-funk"
	"golang.org/x/tools/cover"
)

type funcInfo struct {
	fileName       string
	pkgFileName    string
	functionName   string
	startLine      int
	endLine        int
	uncoveredLines int
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

type FilterOptions struct {
	Include  string
	Exclude  string
	SortFile bool
}

// getProfiles parses profile data in the specified file and returns the parsed profiles.
func getProfiles(filePath string) ([]*cover.Profile, error) {
	profiles, err := cover.ParseProfiles(filePath)
	if err != nil {
		return nil, err
	}

	return profiles, nil
}

// getFunctionInfo returns function information for profiles.
func getFunctionInfo(profiles []*cover.Profile) ([]*funcInfo, int, int, error) {
	total := 0
	covered := 0
	var funcInfos []*funcInfo

	for _, profile := range profiles {
		name := profile.FileName

		filename, err := getFilename(name)
		if err != nil {
			return nil, 0, 0, err
		}

		functions, err := getFunctions(filename)
		if err != nil {
			return nil, 0, 0, err
		}

		for _, f := range functions {
			c, t := f.coverage(profile)

			fi := &funcInfo{
				fileName:       filename,
				pkgFileName:    name,
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

	return funcInfos, covered, total, nil
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

func GetCoverageData(filePath string) ([]*funcInfo, int, int, error) {
	profiles, err := getProfiles(filePath)
	if err != nil {
		return nil, 0, 0, err
	}

	fi, covered, total, err := getFunctionInfo(profiles)
	if err != nil {
		return nil, 0, 0, err
	}

	return fi, covered, total, nil
}

func FilterFunctionInfo(fi []*funcInfo, filterOpts FilterOptions) ([]*funcInfo, error) {
	var filteredFuncInfo []*funcInfo
	var err error

	if filterOpts.Include != "" {
		filteredFuncInfo, err = filterByRegex(filterOpts.Include, fi)
		if err != nil {
			return nil, err
		}
	} else if filterOpts.Exclude != "" {
		filteredFuncInfo, err = filterByRegex(filterOpts.Exclude, fi)
		if err != nil {
			return nil, err
		}

		filteredFuncInfo = funk.Subtract(fi, filteredFuncInfo).([]*funcInfo)
	} else {
		filteredFuncInfo = fi
	}

	if !filterOpts.SortFile {
		fi = sortFuncInfo(filteredFuncInfo)
	} else {
		fi = filteredFuncInfo
	}

	return fi, nil
}
