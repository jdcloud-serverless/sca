package sub_function

import (
	"fmt"
	client2 "github.com/jdcloud-serverless/sca/common/client"
	"github.com/jdcloud-serverless/sca/common/user"
	"github.com/jdcloud-serverless/sca/common/util"

	"github.com/jdcloud-api/jdcloud-sdk-go/services/function/apis"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/function/client"
	"github.com/spf13/cobra"
)

func NewFunctionInfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "get function info in cloud",
		Long:  "get function info in cloud",
		RunE:   infoRun,
	}

	cmd.Flags().StringVarP(&functionName, "name", "n", "", "name of this funtion")
	//cmd.Flags().StringVarP(&version, "version", "v", "", "version of this funtion")
	//cmd.Flags().StringVarP(&alias, "alias", "a", "", "alias of this funtion")
	return cmd
}

func infoRun(cmd *cobra.Command, args []string)error {
	return info(functionName, version, alias)
}

// TODO version and alias for future JSA version
func info(functionName, version, alias string) error{
	user := user.GetUser()
	functionClient := client2.NewFunctionClient(user)

	if functionName == "" {
		return fmt.Errorf("Please input correct function name ...")
	}

	return infoFunction(user, functionClient)

	//if version == "" {
	//	if alias == "" {
	//		// get function
	//		infoFunction(user, functionClient, functionName)
	//	} else {
	//		// get alias
	//		infoAlias(user, functionClient, functionName, alias)
	//	}
	//} else {
	//	if alias == "" {
	//		// get version
	//		infoVersion(user, functionClient, functionName, version)
	//	} else {
	//		fmt.Println("Please choose one version or alias")
	//		return
	//	}
	//}

}

func infoFunction(user *user.User, functionClient *client.FunctionClient) error{
	getFunctionReq := apis.NewGetFunctionRequestWithAllParams(user.Region, functionName)
	getFunctionResp, err := functionClient.GetFunction(getFunctionReq)
	if err != nil || getFunctionResp.Error.Code != 0 {
		if err != nil {
			return fmt.Errorf("Get function (%s) error : %s \n", functionName, err.Error())
		} else {
			return fmt.Errorf("Get function (%s) error : %s \n", functionName, getFunctionResp.Error.Message)
		}
	}
	util.TableFunctionModel(getFunctionResp.Result.Data)
	return nil
}

func infoVersion(user *user.User, functionClient *client.FunctionClient) error {
	getVersionReq := apis.NewGetVersionRequest(user.Region, functionName, version)
	getVersionResp, err := functionClient.GetVersion(getVersionReq)
	if err != nil || getVersionResp.Error.Code != 0 {
		if err != nil {
			return fmt.Errorf("Get function (%s) version (%s) error : %s \n", functionName, version, err.Error())
		} else {
			return fmt.Errorf("Get function (%s) version (%s) error : %s \n", functionName, version, getVersionResp.Error.Message)
		}
	}
	util.TableFunctionModel(getVersionResp.Result.Data)
	return nil
}

func infoAlias(user *user.User, functionClient *client.FunctionClient)error {
	getAliasReq := apis.NewGetAliasRequestWithAllParams(user.Region, functionName, alias)
	getAliasResp, err := functionClient.GetAlias(getAliasReq)
	if err != nil || getAliasResp.Error.Code != 0 {
		if err != nil {
			return fmt.Errorf("Get function (%s) alias (%s) error : %s \n", functionName, alias, err.Error())
		} else {
			return fmt.Errorf("Get function (%s) alias (%s) error : %s \n", functionName, alias, getAliasResp.Error.Message)
		}
	}
	return nil
}
