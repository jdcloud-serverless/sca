package csa

import (
	"csa/commands/build"
	"csa/commands/config"
	"csa/commands/deploy"
	"csa/commands/function"
	"csa/commands/initialize"
	"csa/commands/invoke"
	"csa/commands/local"
	"csa/commands/logs"
	"csa/commands/package"
	"csa/commands/validate"
	"csa/commands/version"

	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Short: "JD Cloud Serverless Application",
}

// add all child command
func init() {
	// version
	RootCommand.AddCommand(version.NewVersionCommand())

	// init
	RootCommand.AddCommand(initialize.NewInitCommand())

	// build
	RootCommand.AddCommand(build.NewBuildCommand())

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

	// package
	RootCommand.AddCommand(_package.NewPackageCommand())

	// validate
	RootCommand.AddCommand(validate.NewValidateCommand())

	// function
	RootCommand.AddCommand(function.NewFunctionCommand())
}

func Start() {
	RootCommand.Execute()
}
