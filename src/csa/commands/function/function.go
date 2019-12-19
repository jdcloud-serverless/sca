package function

import (
	"csa/commands/function/sub_function"

	"github.com/spf13/cobra"
)

func NewFunctionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "function",
		Short: "function",
		Long:  "function",
	}
	cmd.AddCommand(sub_function.NewFunctionListCommand())
	cmd.AddCommand(sub_function.NewFunctionInfoCommand())
	cmd.AddCommand(sub_function.NewFunctionDeleteCommand())
	return cmd
}
