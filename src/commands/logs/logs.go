package logs

import (
	"csa/common"
	"fmt"
	"github.com/spf13/cobra"
	logsApis "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/logs/apis"
	functionApis "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/function/apis"
)

var functionName string
var startTime string
var endTime string
var count int32

func NewLogsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logs",
		Short: "logs",
		Long:  "logs",
		Run:   runLogs,
	}
	cmd.Flags().StringVarP(&functionName, "name", "n", "", "Function name.")
	cmd.Flags().StringVarP(&startTime, "starttime", "s", "", "log start time.")
	cmd.Flags().StringVarP(&endTime, "endtime", "e", "", "log end time.")
	cmd.Flags().Int32VarP(&count, "count", "c", 1000, "count of logs.")

	return cmd
}

func runLogs(cmd *cobra.Command, args []string) {
	if functionName == "" {
		fmt.Printf("function name is empty.")
		return
	}
	// get user info
	user := common.GetUser()
	if logSetId,logTopicId,err :=getFunction(user,functionName);err == nil{
		findLog(user,logSetId,logTopicId)
	}
}

func getFunction(user*common.User,functionName string)  (logSetId,logTopicId string,err error){
	req := functionApis.NewGetFunctionRequestWithAllParams(user.Region, functionName)
	resp, err := common.NewFunctionClient(user).GetFunction(req)
	if err != nil || resp.Error.Code != 0 {
		if err != nil {
			fmt.Printf("Get function (%s) error : %s \n", functionName, err.Error())
		} else {
			fmt.Printf("Get function (%s) error : %s \n", functionName, resp.Error.Message)
		}
		return "","",err
	}
	return resp.Result.Data.LogSetId,resp.Result.Data.LogTopicId,nil
}

func findLog(user*common.User,logSetId,logTopicId string) {
	// https://docs.jdcloud.com/cn/log-service/api/search?content=API
	client := common.NewLogClient(user)
	req := logsApis.NewSearchRequest(user.Region, logSetId, logTopicId, "fulltext")
	req.SetPageSize(int(count))
	req.SetPageNumber(1)
	if startTime != "" {
		req.SetStartTime(startTime)
	}
	if endTime != "" {
		req.SetStartTime(endTime)
	}

	resp, err := client.Search(req)
	if err != nil {
		fmt.Printf("find log err=%s\n", err.Error())
		return
	}
	if resp.Error.Code != 0 || resp.Error.Code != 200 {
		fmt.Printf("find log err=%s\n", resp.Error.Message)
		return
	}
	fmt.Printf("logs:\n%v\n", resp.Result.Data)
}
