package sub_function

import (
	"fmt"
	local_client "github.com/jdcloud-serverless/sca/common/client"
	"github.com/jdcloud-serverless/sca/common/user"

	"github.com/jdcloud-api/jdcloud-sdk-go/services/function/apis"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/function/client"
	"github.com/spf13/cobra"
)

func NewFunctionDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete function in cloud",
		Long:  "delete function in cloud",
		RunE:   deleteRun,
	}
	cmd.Flags().StringVarP(&functionName, "name", "n", "", "name of this funtion")
	//cmd.Flags().StringVarP(&version, "version", "v", "", "version of this funtion")
	//cmd.Flags().StringVarP(&alias, "alias", "a", "", "alias of this funtion")
	return cmd
}

func deleteRun(cmd *cobra.Command, args []string) error{
	return delete(functionName, version, alias)
}

// TODO version and alias for future JSA version
func delete(functionName, version, alias string) error {
	user := user.GetUser()
	functionClient := local_client.NewFunctionClient(user)

	if functionName == "" {
		return fmt.Errorf("Please input correct function name ...")
	}

	return deleteFunction(user, functionClient)

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

func deleteFunction(user *user.User, functionClient *client.FunctionClient) error {
	deleteFunctionReq := apis.NewDeleteFunctionRequestWithAllParams(user.Region, functionName)
	deleteFunctionResp, err := functionClient.DeleteFunction(deleteFunctionReq)
	if err != nil || deleteFunctionResp.Error.Code != 0 {
		if err != nil {
			return fmt.Errorf("Delete function (%s) error : %s \n", functionName, err.Error())
		} else {
			return fmt.Errorf("Delete function (%s) error : %s \n", functionName, deleteFunctionResp.Error.Message)
		}
	}
	return nil
}

func deleteVersion(user *user.User, functionClient *client.FunctionClient) error{
	deleteVersionReq := apis.NewDeleteVersionRequestWithAllParams(user.Region, functionName, version)
	deleteVersionResp, err := functionClient.DeleteVersion(deleteVersionReq)
	if err != nil || deleteVersionResp.Error.Code != 0 {
		if err != nil {
			return fmt.Errorf("Delete function (%s) version (%s) error : %s \n", functionName, version, err.Error())
		} else {
			return fmt.Errorf("Delete function (%s) version (%s) error : %s \n", functionName, version, deleteVersionResp.Error.Message)
		}
	}
	return nil
}

func deleteAlias(user *user.User, functionClient *client.FunctionClient)error {
	deleteAliasReq := apis.NewDeleteAliasRequestWithAllParams(user.Region, functionName, alias)
	deleteAliasResp, err := functionClient.DeleteAlias(deleteAliasReq)
	if err != nil || deleteAliasResp.Error.Code != 0 {
		if err != nil {
			return fmt.Errorf("Delete function (%s) alias (%s) error : %s \n", functionName, alias, err.Error())
		} else {
			return fmt.Errorf("Delete function (%s) alias (%s) error : %s \n", functionName, alias, deleteAliasResp.Error.Message)
		}
	}
	return nil
}
