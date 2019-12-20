package main

import (
	"github.com/jdcloud-serverless/sca/commands/config"
	"github.com/jdcloud-serverless/sca/commands/deploy"
	"github.com/jdcloud-serverless/sca/commands/function"
	"github.com/jdcloud-serverless/sca/commands/initialize"
	"github.com/jdcloud-serverless/sca/commands/invoke"
	"github.com/jdcloud-serverless/sca/commands/local"
	"github.com/jdcloud-serverless/sca/commands/logs"
	"github.com/jdcloud-serverless/sca/commands/validate"
	"github.com/jdcloud-serverless/sca/commands/version"

	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Short: "JD Serverless Cloud Application",
}

// add all child command
func init() {
	// version
	RootCommand.AddCommand(version.NewVersionCommand())

	// init
	RootCommand.AddCommand(initialize.NewInitCommand())

	// build
	//RootCommand.AddCommand(build.NewBuildCommand())

	// config
	RootCommand.AddCommand(config.NewConfigCommand())

	// deploy
	RootCommand.AddCommand(deploy.NewDeployCommand())

	// invoke
	RootCommand.AddCommand(invoke.NewInvokeCommand())

	// local
	RootCommand.AddCommand(local.NewLocalCommand())

	// logs
	RootCommand.AddCommand(logs.NewLogsCommand())

	// validate
	RootCommand.AddCommand(validate.NewValidateCommand())

	// function
	RootCommand.AddCommand(function.NewFunctionCommand())
}

func main() {
	RootCommand.Execute()
}
