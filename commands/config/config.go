package config

import (
	"fmt"
	"github.com/jdcloud-serverless/sca/common/user"
	"os"

	"github.com/spf13/cobra"
)

const ConfigFileTemplate = `account_id=%s
region=%s
access_key=%s
secret_key=%s
`

func NewConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "config user credential info",
		Long:  "config user credential info",
		Run:   config,
	}
	return cmd
}

func config(cmd *cobra.Command, args []string) {
	fmt.Print("[>] JDCould accountid = ")
	var accountId string
	fmt.Scan(&accountId)

	fmt.Print("[>] JDCould region = ")
	var region string
	fmt.Scan(&region)

	fmt.Print("[>] JDCould access-key = ")
	var access_key string
	fmt.Scan(&access_key)

	fmt.Print("[>] JDCould secret-key = ")
	var secret_key string
	fmt.Scan(&secret_key)

	setConfigFile(accountId, region, access_key, secret_key)
}

func setConfigFile(accountId, region, access_key, secret_key string) {
	fileString := fmt.Sprintf(ConfigFileTemplate, accountId, region, access_key, secret_key)

	f, err := os.OpenFile(user.UserInfoPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
	} else {
		// offset
		//os.Truncate(filename, 0) //clear
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(fileString), n)
		defer f.Close()
	}
}
