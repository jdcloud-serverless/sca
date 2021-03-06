package invoke

import (
	"fmt"
	local_client "github.com/jdcloud-serverless/sca/common/client"
	"github.com/jdcloud-serverless/sca/common/user"
	"io/ioutil"
	"os"

	"github.com/jdcloud-api/jdcloud-sdk-go/services/function/apis"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/function/client"
	"github.com/spf13/cobra"
)

const (
	FunctionLatestVersion = "LATEST"
	InvokeResultFormat    = "RequestId: %s \t\t Billed Duration: %d ms \t\tMemory Size: %d MB \t\tMax Memory Used : %.2f MB\n"
	InvokeReturnFormat    = "Invoke Return : %s\n"
)

var (
	functionName, version, alias, eventFile string

	async bool
)

func NewInvokeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invoke",
		Short: "invoke function in cloud",
		Long:  "invoke function in cloud",
		RunE:   invoke,
	}
	InitInvokeCmdFlags(cmd)
	return cmd
}

func InitInvokeCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&functionName, "name", "n", "", "specify function name to invoke")
	//cmd.Flags().StringVarP(&version, "version", "v", "", "specify function version to invoke")
	//cmd.Flags().StringVarP(&alias, "alias", "a", "", "specify function alias to invoke")
	cmd.Flags().StringVarP(&eventFile, "event", "e", "", `specify event file path to invoke, if not specified, use event {"key":"value"}`)
	//cmd.Flags().BoolVarP(&async, "async", "a", false, "specify invoke method to async")
}

func invoke(cmd *cobra.Command, args []string) error {
	if functionName == "" {
		fmt.Println("Invoke Error : Please input correct function name")
	}

	user := user.GetUser()
	functionClient := local_client.NewFunctionClient(user)

	eventStr := ""
	if eventFile == ""{
		eventStr = `{"key":"value"}`
	} else {
		var err error = nil
		eventStr, err = readEventFile(eventFile)
		if err != nil {
			return fmt.Errorf("read event file err=",err)
		}
	}

	if async {
		return asyncInvokeFunction(user, functionClient, eventStr)
	} else {
		return syncInvokeFunction(user, functionClient, eventStr)
	}
}

func readEventFile(eventFile string) (string, error) {
	var realPath string
	if eventFile[0] == '.' {
		curDir, _ := os.Getwd()
		realPath = curDir + eventFile[1:]
	} else {
		realPath = eventFile
	}

	eventByte, _ := ioutil.ReadFile(realPath)
	return string(eventByte), nil
}

func syncInvokeFunction(user *user.User, client *client.FunctionClient, eventStr string) error {
	invokeReq := apis.NewInvokeRequestWithAllParams(user.Region, functionName, FunctionLatestVersion, eventStr)
	invokeResp, err := client.Invoke(invokeReq)
	if err != nil || invokeResp.Error.Code != 0 {
		if err == nil {
			return fmt.Errorf("Invoke function (%s) error : \n %s\n", functionName, invokeResp.Error.Message)
		} else {
			return fmt.Errorf("Invoke function (%s) error : \n %s\n", functionName, err.Error())
		}
	}
	fmt.Println(invokeResp.Result.Data.LogStr)
	fmt.Printf(InvokeReturnFormat, invokeResp.Result.Data.Result)
	fmt.Printf(InvokeResultFormat, invokeResp.RequestID, invokeResp.Result.Data.BillingTime, invokeResp.Result.Data.SetupMem, invokeResp.Result.Data.RealMem)
	return nil
}

func asyncInvokeFunction(user *user.User, client *client.FunctionClient, eventStr string)  error{
	asyncInvokeReq := apis.NewAsyncInvokeRequestWithAllParams(user.Region, functionName, FunctionLatestVersion, eventStr)
	asyncInvokeResp, err := client.AsyncInvoke(asyncInvokeReq)
	if err != nil || asyncInvokeResp.Error.Code == 0 {
		if err == nil {
			return fmt.Errorf("Invoke function (%s) error : \n %s\n", functionName, asyncInvokeResp.Error.Message)
		} else {
			return fmt.Errorf("Invoke function (%s) error : \n %s\n", functionName, err.Error())
		}
	}
	return nil
}
