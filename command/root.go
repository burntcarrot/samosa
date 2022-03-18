package command

import (
	"log"

	"github.com/burntcarrot/samosa/internal"
	"github.com/spf13/cobra"
)

type Options struct {
	File string
	Pkg  bool
}

func NewCmdRoot() *cobra.Command {
	opts := Options{}
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVarP(&opts.File, "file", "f", "", "Coverage file path")
	cmd.Flags().BoolVarP(&opts.Pkg, "pkg", "p", false, "Use package-based path")

	return cmd
}

func (opts *Options) Run() error {
	if opts.File != "" {
		fi, covered, total, err := internal.GetCoverageData(opts.File)
		if err != nil {
			log.Fatalf("failed to get coverage data: %v\n", err)
		}

		internal.PrintTable(fi, covered, total, opts.Pkg)
	}
	return nil
}

func Execute() error {
	cmd := NewCmdRoot()
	return cmd.Execute()
}
