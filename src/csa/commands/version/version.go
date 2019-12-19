package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	Version = "0.0.1"
)

func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "version",
		Long:  "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("JD Cloud Serverless Application Version:", Version)
		},
	}
	return cmd
}
