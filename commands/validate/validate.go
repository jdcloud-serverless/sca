package validate

import (
	"fmt"

	"github.com/jdcloud-serverless/sca/common"

	"github.com/spf13/cobra"
)

var templateFileName string

func NewValidateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "validate",
		Long:  "validate",
		Run:   runValidate,
	}
	cmd.Flags().StringVarP(&templateFileName, "template-file", "t", "", "The template file.")
	return cmd
}

func runValidate(cmd *cobra.Command, args []string) {
	if templateFileName == "" {
		fmt.Println("please input template file name.")
		return
	}
	template, err := common.LoadTemplate(templateFileName)
	if err != nil {
		return
	}

	for functionName, tmpl := range template.Resources {
		if !FunctionNameCheck(functionName) {
			fmt.Printf("FunctionName(%s) is invalid.", functionName)
			return
		}
		if !RuntimeCheck(tmpl.FunctionProperties.Runtime) {
			fmt.Printf("Runtime(%s) is not support.", tmpl.FunctionProperties.Runtime)
			return
		}
		if !HandlerCheck(tmpl.FunctionProperties.Runtime, tmpl.FunctionProperties.Handler) {
			fmt.Printf("Handler(%s) is not invalid.", tmpl.FunctionProperties.Handler)
			return
		}

		if !MemoryCheck(tmpl.FunctionProperties.MemorySize) {
			fmt.Printf("MemorySize(%d) is not support.", tmpl.FunctionProperties.MemorySize)
			return
		}
		if !OvertimeCheck(tmpl.FunctionProperties.Timeout) {
			fmt.Printf("Timeout(%d) is not support.", tmpl.FunctionProperties.Timeout)
			return
		}

		if !EnvCheck(tmpl.FunctionProperties.Env) {
			fmt.Printf("Env(%v) is invalid.", tmpl.FunctionProperties.Env)
			return
		}

		if !CodeUriCheck(tmpl.FunctionProperties.CodeUri) {
			fmt.Printf("CodeUri(%s) is invalid.", tmpl.FunctionProperties.CodeUri)
			return
		}
	}
}
