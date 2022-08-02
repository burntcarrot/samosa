package samosa

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
)

type FilterOptions struct {
	Include  string
	Exclude  string
	SortFile bool
}

type Options struct {
	File          string
	Export        string
	OutputFile    string
	Pkg           bool
	SortFile      bool
	FilterOptions FilterOptions
}

func NewCmdRoot() *cobra.Command {
	opts := Options{}
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVarP(&opts.File, "file", "f", "", "Coverage file path")
	cmd.Flags().StringVarP(&opts.FilterOptions.Include, "include", "i", "", "Include results for specified file")
	cmd.Flags().StringVarP(&opts.FilterOptions.Exclude, "exclude", "x", "", "Exclude results for specified file")
	cmd.Flags().StringVarP(&opts.OutputFile, "output", "o", "", "Output filename for exporting results")
	cmd.Flags().StringVarP(&opts.Export, "export", "e", "", "Export results to CSV, JSON, etc.")
	cmd.Flags().BoolVarP(&opts.Pkg, "pkg", "p", false, "Use package-based path")
	cmd.Flags().BoolVarP(&opts.SortFile, "sort-file", "s", false, "Sort results based on file path")

	return cmd
}

func (opts *Options) Run() error {
	fi, covered, total, err := GetCoverageData(opts.File)
	if err != nil {
		log.Default().Fatalf("failed to get coverage data: %v\n", err)
	}

	filterOpts := FilterOptions{
		Include:  opts.FilterOptions.Include,
		Exclude:  opts.FilterOptions.Exclude,
		SortFile: opts.SortFile,
	}

	fi, err = FilterFunctionInfo(fi, filterOpts)
	if err != nil {
		log.Fatalf("failed to get function info: %v\n", err)
	}

	if opts.File != "" {
		switch strings.TrimSpace(opts.Export) {
		case "json":
			err = ExportJSON(opts.OutputFile, fi)
		case "csv":
			err = ExportCSV(opts.OutputFile, fi)
		default:
			PrintTable(fi, covered, total, opts.Pkg)
		}
	}

	if err != nil {
		log.Fatalf("failed to export results: %v\n", err)
	}

	return nil
}

func Execute() error {
	cmd := NewCmdRoot()
	return cmd.Execute()
}
