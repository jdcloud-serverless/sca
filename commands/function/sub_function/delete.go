package sub_function

import (
	"fmt"

	"github.com/jdcloud-serverless/sca/common"

	"github.com/jdcloud-api/jdcloud-sdk-go/services/function/apis"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/function/client"
	"github.com/spf13/cobra"
)

func NewFunctionDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete",
		Long:  "delete",
		Run:   deleteRun,
	}
	cmd.Flags().StringVarP(&functionName, "name", "n", "", "name of this funtion")
	//cmd.Flags().StringVarP(&version, "version", "v", "", "version of this funtion")
	//cmd.Flags().StringVarP(&alias, "alias", "a", "", "alias of this funtion")
	return cmd
}

func deleteRun(cmd *cobra.Command, args []string) {
	delete(functionName, version, alias)
}

// TODO version and alias for future JSA version
func delete(functionName, version, alias string) {
	user := common.GetUser()
	functionClient := common.NewFunctionClient(user)

	if functionName == "" {
		fmt.Println("Please input correct function name ...")
		return
	}

	deleteFunction(user, functionClient)

	//if version == "" {
	//	if alias == "" {
	//		// invoke delete function
	//		deleteFunction(user, functionClient, functionName)
	//	} else {
	//		// invoke delete alias
	//		deleteAlias(user, functionClient, functionName, alias)
	//	}
	//} else {
	//	if alias == "" {
	//		// invoke delete version
	//		deleteVersion(user, functionClient, functionName, version)
	//	} else {
	//		fmt.Println("Please choose one version or alias")
	//		return
	//	}
	//}
}

func deleteFunction(user *common.User, functionClient *client.FunctionClient) {
	deleteFunctionReq := apis.NewDeleteFunctionRequestWithAllParams(user.Region, functionName)
	deleteFunctionResp, err := functionClient.DeleteFunction(deleteFunctionReq)
	if err != nil || deleteFunctionResp.Error.Code != 0 {
		if err != nil {
			fmt.Printf("Delete function (%s) error : %s \n", functionName, err.Error())
		} else {
			fmt.Printf("Delete function (%s) error : %s \n", functionName, deleteFunctionResp.Error.Message)
		}
	}
}

func deleteVersion(user *common.User, functionClient *client.FunctionClient) {
	deleteVersionReq := apis.NewDeleteVersionRequestWithAllParams(user.Region, functionName, version)
	deleteVersionResp, err := functionClient.DeleteVersion(deleteVersionReq)
	if err != nil || deleteVersionResp.Error.Code != 0 {
		if err != nil {
			fmt.Printf("Delete function (%s) version (%s) error : %s \n", functionName, version, err.Error())
		} else {
			fmt.Printf("Delete function (%s) version (%s) error : %s \n", functionName, version, deleteVersionResp.Error.Message)
		}
	}
}

func deleteAlias(user *common.User, functionClient *client.FunctionClient) {
	deleteAliasReq := apis.NewDeleteAliasRequestWithAllParams(user.Region, functionName, alias)
	deleteAliasResp, err := functionClient.DeleteAlias(deleteAliasReq)
	if err != nil || deleteAliasResp.Error.Code != 0 {
		if err != nil {
			fmt.Printf("Delete function (%s) alias (%s) error : %s \n", functionName, alias, err.Error())
		} else {
			fmt.Printf("Delete function (%s) alias (%s) error : %s \n", functionName, alias, deleteAliasResp.Error.Message)
		}
	}
}
