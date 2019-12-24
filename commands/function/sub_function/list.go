package sub_function

import (
	"errors"
	"fmt"
	client2 "github.com/jdcloud-serverless/sca/common/client"
	"github.com/jdcloud-serverless/sca/common/user"
	"os"

	"github.com/jdcloud-api/jdcloud-sdk-go/services/function/apis"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/function/client"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func NewFunctionListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "get list of functions in cloud",
		Long:  "get list of functions in cloud",
		Run: func(cmd *cobra.Command, args []string) {
			list()
		},
	}
	return cmd
}

func list() {
	user := user.GetUser()
	client := client2.NewFunctionClient(user)

	listResp, err := listFunction(user, client)
	if err != nil {
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Function Name", "Description", "Runtime", "Code url", "Create Time"})
	for _, v := range listResp.Result.Data.Functions {
		row := []string{v.Name, v.Description, v.Runtime, v.DownloadUrl, v.CreateTime}
		table.Append(row)
	}
	table.Render()
}

func listFunction(user *user.User, client *client.FunctionClient) (*apis.ListFunctionResponse, error) {
	listReq := apis.NewListFunctionRequestWithoutParam()
	listReq.SetRegionId(user.Region)
	listReq.SetListAll(true)

	listResp, err := client.ListFunction(listReq)
	if err != nil || listResp.Error.Code != 0 {
		if err == nil {
			fmt.Printf("List functions error : \n %s\n", listResp.Error.Message)
			return nil, err
		} else {
			fmt.Printf("List functions error : \n %s\n", err.Error())
			return nil, errors.New(err.Error())
		}
	}
	return listResp, nil
}
