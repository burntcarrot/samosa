package command

import (
	"log"
	"strings"

	"github.com/burntcarrot/samosa/internal"
	"github.com/spf13/cobra"
)

type Options struct {
	File          string
	Export        string
	OutputFile    string
	Pkg           bool
	SortFile      bool
	FilterOptions FilterOptions
}

type FilterOptions struct {
	Include string
	Exclude string
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
	fi, covered, total, err := internal.GetCoverageData(opts.File)
	if err != nil {
		log.Fatalf("failed to get coverage data: %v\n", err)
	}

	filterOpts := internal.FilterOptions{
		Include:  opts.FilterOptions.Include,
		Exclude:  opts.FilterOptions.Exclude,
		SortFile: opts.SortFile,
	}

	fi, err = internal.FilterFunctionInfo(fi, filterOpts)
	if err != nil {
		log.Fatalf("failed to get function info: %v\n", err)
	}

	if opts.File != "" {
		switch strings.TrimSpace(opts.Export) {
		case "json":
			internal.ExportJSON(opts.OutputFile, fi)
		case "csv":
			internal.ExportCSV(opts.OutputFile, fi)
		default:
			internal.PrintTable(fi, covered, total, opts.Pkg)
		}
	}
	return nil
}

func Execute() error {
	cmd := NewCmdRoot()
	return cmd.Execute()
}
