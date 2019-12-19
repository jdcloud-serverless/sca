package build

import (
	"github.com/spf13/cobra"
)

func NewBuildCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "build",
		Long:  "build",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd
}
