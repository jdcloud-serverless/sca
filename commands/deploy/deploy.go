package deploy

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/function/apis"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/function/client"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/function/models"
	"github.com/jdcloud-serverless/sca/common"
	"github.com/mholt/archiver"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
)

const ZipFileSuffix = "code.zip"

var templatePath string

func NewDeployCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "deploy functions in template to cloud",
		Long:  "deploy functions in template to cloud",
		Run:   deploy,
	}
	InitDeployCmdFlags(cmd)
	return cmd
}

func InitDeployCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&templatePath, "template", "t", "", "specify template yaml file")
}

func deploy(cmd *cobra.Command, args []string) {
	// get template

	template, err := common.LoadTemplate(templatePath)
	if err != nil {
		return
	}

	// get user info
	user := common.GetUser()

	for functionName, v := range template.Resources {
		// create function sdk client
		functionClient := common.NewFunctionClient(user)

		// if function not exists , execute create function
		// if exists , execute update function
		if exist := checkFunctionExist(user.Region, functionName, functionClient); !exist {
			fmt.Printf("Function (%s) not exists , now beginning to create function\n", functionName)
			createFunction(functionName, user.Region, &v.FunctionProperties, functionClient)
		} else {
			fmt.Printf("Function (%s) exists , now beginning to update function\n", functionName)
			updateFunction(functionName, user.Region, &v.FunctionProperties, functionClient)
		}
	}

}

func checkFunctionExist(region, functionName string, functionClient *client.FunctionClient) bool {
	getFunctionReq := apis.NewGetFunctionRequestWithAllParams(region, functionName)
	getFunctionResp, err := functionClient.GetFunction(getFunctionReq)

	if err != nil || getFunctionResp.Error.Code != 0 {
		return false
	}
	return true
}

func createFunction(functionName, region string, function *common.FunctionProperties, functionClient *client.FunctionClient) {
	createFunctionReq := apis.NewCreateFunctionRequestWithoutParam()
	createFunctionReq.SetRegionId(region)
	createFunctionReq.SetName(functionName)
	createFunctionReq.SetDescription(function.Description)
	createFunctionReq.SetEntrance(function.Handler)
	createFunctionReq.SetEnvironment(getEnvRequestByEnv(function.Env))
	createFunctionReq.SetOverTime(function.Timeout)
	createFunctionReq.SetMemory(function.MemorySize)
	createFunctionReq.SetRunTime(function.Runtime)
	codeReq, err := getCodeRequestByCodeUri(function.CodeUri)
	if err != nil {
		fmt.Printf("Deploy function (%s) error !\n", functionName)
		return
	}
	createFunctionReq.SetCode(codeReq)

	createFunctionResp, err := functionClient.CreateFunction(createFunctionReq)
	if err != nil || createFunctionResp.Error.Code != 0 {
		if err == nil {
			fmt.Printf("Deploy function (%s) error : \n %s\n", functionName, createFunctionResp.Error.Message)
		} else {
			fmt.Printf("Deploy function (%s) error : \n %s\n", functionName, err.Error())
		}
		return
	}

	common.TableFunctionModel(createFunctionResp.Result.Data)
	fmt.Printf("Deploy function (%s) success .\n", functionName)
}

func updateFunction(functionName, region string, function *common.FunctionProperties, functionClient *client.FunctionClient) {
	updateFunctionReq := apis.NewUpdateFunctionRequestWithoutParam()
	updateFunctionReq.SetRegionId((region))
	updateFunctionReq.SetFunctionName(functionName)
	updateFunctionReq.SetRunTime(function.Runtime)
	updateFunctionReq.SetDescription(function.Description)
	updateFunctionReq.SetEntrance(function.Handler)
	updateFunctionReq.SetEnvironment(getEnvRequestByEnv(function.Env))
	updateFunctionReq.SetOverTime(function.Timeout)
	updateFunctionReq.SetMemory(function.MemorySize)
	codeReq, err := getCodeRequestByCodeUri(function.CodeUri)
	if err != nil {
		fmt.Printf("Deploy function (%s) error !\n%s\n", functionName, err.Error())
		return
	}
	updateFunctionReq.SetCode(codeReq)

	updateFunctionResp, err := functionClient.UpdateFunction(updateFunctionReq)
	if err != nil || updateFunctionResp.Error.Code != 0 {
		if err == nil {
			fmt.Printf("Deploy function (%s) error : \n %s\n", functionName, updateFunctionResp.Error.Message)
		} else {
			fmt.Printf("Deploy function (%s) error : \n %s\n", functionName, err.Error())
		}
		return
	}

	common.TableFunctionModel(updateFunctionResp.Result.Data)
	fmt.Printf("Deploy function (%s) success .\n", functionName)
}

func getCodeRequestByCodeUri(codeUri string) (*models.Code, error) {
	codeZipUri, err := compress(codeUri)
	if err != nil {
		return nil, err
	}

	file, err := readCodeFile(codeZipUri)
	if err != nil {
		return nil, err
	}
	os.Remove(codeZipUri)

	zipfile := base64.StdEncoding.EncodeToString(file)

	return &models.Code{
		ZipFile: &zipfile,
	}, nil
}

func getCodeRealPath(codeUri string) (string, error) {
	if len(codeUri) == 0 {
		return "", errors.New("[Error] Code Uri is empty ...")
	}
	templateAbsDir := getFileAbsDir(templatePath)
	if codeUri[0] == '.' {
		return templateAbsDir+codeUri[1:], nil
	}  else if codeUri[0] == '/'{
		return codeUri,nil
	} else{
		return templateAbsDir + "/" + codeUri, nil
	}
}

func getFileAbsDir(path string)string{
	res := ""
	if path[0] == '.' {
		curDir,_ := os.Getwd()
		res =  curDir+path[1:]
	}  else if path[0] == '/'{
		res =  path
	} else{
		curDir,_ := os.Getwd()
		res =  curDir + "/" + path
	}
	lastIndex := strings.LastIndex(res,"/")
	return res[:lastIndex]
}

func compress(codeUri string) (string, error) {
	dirPath, err := getCodeRealPath(codeUri)
	if err != nil {
		return "", err
	}
	zipFilePath := dirPath + ZipFileSuffix
	err = archiver.Archive([]string{dirPath}, zipFilePath)
	if err != nil {
		return "", err
	}
	return zipFilePath, nil
}

func readCodeFile(codeUri string) ([]byte, error) {
	b, err := ioutil.ReadFile(codeUri)
	if err != nil {
		fmt.Printf("Open code zip file error : %s\n", err.Error())
		return nil, err
	}
	return b, nil
}

func getEnvRequestByEnv(env map[string]string) *models.Env {
	var a interface{}
	a = env
	res := &models.Env{&a}
	return res
}
