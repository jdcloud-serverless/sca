package local

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"csa/common"

	"github.com/spf13/cobra"
)

const (
	WsgiUrl = "http://127.0.0.1:9090/invoke"
)

func NewLocalCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "local [flags] [function_name]",
		Short: "invoke a function in local container",
		Long:  "invoke a function in local container",
		RunE:  ExecLocalCommand,
	}

	InitLocalCmdFlags(cmd)

	return cmd
}

func InitLocalCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("template", "t", "", "specify template yaml file")
	cmd.Flags().StringP("event", "e", "", "specify event json file")
}

func ExecLocalCommand(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("function handler missing")
	}
	functionName := args[0]

	tFile, err := cmd.Flags().GetString("template")
	if err != nil {
		return err
	}

	eFile, err := cmd.Flags().GetString("event")
	if err != nil {
		return err
	}

	template, err := common.LoadTemplate(tFile)
	if err != nil {
		return err
	}

	event, err := common.ReadFile(eFile)
	if err != nil {
		return err
	}

	if _, ok := template.Resources[functionName]; !ok {
		return fmt.Errorf("Not find function name(%s)", functionName)
	}
	funcProperties := template.Resources[functionName].FunctionProperties

	//start docker container
	dCliet, err := NewDockerClient()
	if err != nil {
		return err
	}
	if err := dCliet.PullImage(convertRuntime(funcProperties.Runtime)); err != nil {
		return err
	}
	codePath, err := getCodeAbsPath(tFile, funcProperties.CodeUri)
	if err != nil {
		return err
	}
	if err := dCliet.CreateDocker(codePath, convertRuntime(funcProperties.Runtime)); err != nil {
		return err
	}
	if err := dCliet.StartDocker(); err != nil {
		return err
	}
	defer dCliet.RemoveDocker()

	//send http request
	time.Sleep(time.Second)
	resp, err := json.MarshalIndent(Execute(functionName, funcProperties, event), "", "	")
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", string(resp))

	return nil
}

type LocalFunctionResponseMessage struct {
	Code       int    `json:"code"`
	Return     string `json:"return"`
	Stdout     string `json:"stdout"`
	Stderr     string `json:"stderr"`
	MemoryUsed string `json:"memory_used"`
	Duration   string `json:"time_used"`
}

func Execute(functionName string, properties common.FunctionProperties, event []byte) *LocalFunctionResponseMessage {
	res := new(LocalFunctionResponseMessage)

	requestId := fmt.Sprintf("%s-%d", "csa-requestid", time.Now().Unix())
	functionId := fmt.Sprintf("%s-%d", "csa-functionid", time.Now().Unix())

	startTime := time.Now()
	header := make(map[string]string)
	header[HeaderContentType] = TypeJson
	header[HeaderRequestId] = requestId
	header[HeaderFunctionId] = functionId
	header[HeaderFunctionName] = functionName
	header[HeaderFunctionVersion] = "csa-local-v1.0"
	header[HeaderFunctionHandler] = properties.Handler
	header[HeaderFunctionMemory] = strconv.Itoa(properties.MemorySize)
	header[HeaderFunctionTimeout] = strconv.Itoa(properties.Timeout)
	header[HeaderFunctionMD5] = fmt.Sprintf("%s-%d", "csa-md5", time.Now().Unix())
	header[HeaderFunctionCodePath] = TARGET_VOL_PATH
	env, _ := json.Marshal(properties.Env)
	header[HeaderFunctionEnv] = string(env)

	postData, err := json.Marshal(event)
	if err != nil {
		res.Code = InternalErrorCode
		return res
	}
	wsgiClient := common.NewHttpClient()
	httpRsp, err := wsgiClient.Forward(WsgiUrl, http.MethodPost, bytes.NewReader(postData), header)
	if err != nil {
		res.Stderr = err.Error()
		res.Code = InvokeFunctionAskCode
		return res
	}

	costTime := time.Since(startTime)
	res.Duration = costTime.String()

	wsgiRes := new(WsgiResponse)
	data, err := ioutil.ReadAll(httpRsp.Body)
	httpRsp.Body.Close()
	if err != nil {
		res.Stderr = err.Error()
		res.Code = InvokeFunctionAskCode
		return res
	}
	if err = json.Unmarshal(data, wsgiRes); err != nil {
		res.Stderr = err.Error()
		res.Code = InvokeFunctionAskCode
		return res
	}

	res.Return = wsgiRes.FunctionReturn
	res.MemoryUsed = fmt.Sprintf("%.2fm", wsgiRes.MemUsage)
	res.Stdout = wsgiRes.StdoutReturn
	res.Stderr = wsgiRes.StderrReturn

	if httpRsp.StatusCode != 200 {
		if httpRsp.StatusCode == 400 {
			res.Code = UserFuncExecuteError
		} else {
			res.Code = InvokeFunctionAskCode
		}
		return res
	}

	res.Code = InternalSuccessCode
	return res
}

func convertRuntime(runtime string) string {
	localRuntime := ""
	switch runtime {
	case common.RUNTIME_Python2_7:
		localRuntime = RUNTIME_Python27
	case common.RUNTIME_Python3_6:
		localRuntime = RUNTIME_Python36
	case common.RUNTIME_Python3_7:
		localRuntime = RUNTIME_Python37
	}

	return localRuntime
}

func getCodeAbsPath(templateFile, codePath string) (string, error) {
	templatePath, err := filepath.Abs(filepath.Dir(templateFile))
	if err != nil {
		return "", err
	}

	oldPwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if err := os.Chdir(templatePath); err != nil {
		return "", err
	}
	defer os.Chdir(oldPwd)

	path, err := filepath.Abs(codePath)
	if err != nil {
		return "", err
	}

	return path, nil
}
