package validate

import (
	"fmt"
	"github.com/jdcloud-serverless/sca/common/template"

	"github.com/spf13/cobra"
)

var templateFileName string

func NewValidateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "validate template",
		Long:  "validate template",
		RunE:   runValidate,
	}
	cmd.Flags().StringVarP(&templateFileName, "template", "t", "", "The template file.")
	return cmd
}

func runValidate(cmd *cobra.Command, args []string) error{
	if templateFileName == "" {
		return fmt.Errorf("please input template file name.")
	}
	template, err := template.LoadTemplate(templateFileName)
	if err != nil {
		return err
	}

	for functionName, tmpl := range template.Resources {
		if !FunctionNameCheck(functionName) {
			return fmt.Errorf("FunctionName(%s) is invalid.", functionName)
		}
		if !RuntimeCheck(tmpl.FunctionProperties.Runtime) {
			return fmt.Errorf("Runtime(%s) is not support.", tmpl.FunctionProperties.Runtime)
		}
		if !HandlerCheck(tmpl.FunctionProperties.Runtime, tmpl.FunctionProperties.Handler) {
			return fmt.Errorf("Handler(%s) is not invalid.", tmpl.FunctionProperties.Handler)
		}

		if !MemoryCheck(tmpl.FunctionProperties.MemorySize) {
			return fmt.Errorf("MemorySize(%d) is not support.", tmpl.FunctionProperties.MemorySize)
		}
		if !OvertimeCheck(tmpl.FunctionProperties.Timeout) {
			return fmt.Errorf("Timeout(%d) is not support.", tmpl.FunctionProperties.Timeout)
		}

		if !EnvCheck(tmpl.FunctionProperties.Env) {
			return fmt.Errorf("Env(%v) is invalid.", tmpl.FunctionProperties.Env)
		}

		if !CodeUriCheck(tmpl.FunctionProperties.CodeUri) {
			return fmt.Errorf("CodeUri(%s) is invalid.", tmpl.FunctionProperties.CodeUri)
		}
	}
	fmt.Println("validate success.")
	return nil
}
