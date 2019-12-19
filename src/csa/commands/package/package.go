package _package

import (
	"github.com/spf13/cobra"
)

func NewPackageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "package",
		Short: "package",
		Long:  "package",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd
}
