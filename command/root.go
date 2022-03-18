package command

import (
	"log"

	"github.com/burntcarrot/samosa/internal"
	"github.com/spf13/cobra"
)

type Options struct {
	File string
}

func NewCmdRoot() *cobra.Command {
	opts := Options{}
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVarP(&opts.File, "file", "f", "", "Coverage file path")

	return cmd
}

func (opts *Options) Run() error {
	if opts.File != "" {
		err := internal.GetCoverageData(opts.File, true)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return nil
}

func Execute() error {
	cmd := NewCmdRoot()
	return cmd.Execute()
}
