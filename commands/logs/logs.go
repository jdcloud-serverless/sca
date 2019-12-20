package logs

import (
	"encoding/json"
	"fmt"

	"github.com/jdcloud-serverless/sca/common"

	functionApis "github.com/jdcloud-api/jdcloud-sdk-go/services/function/apis"
	logsApis "github.com/jdcloud-api/jdcloud-sdk-go/services/logs/apis"
	"github.com/spf13/cobra"
	"time"
)

var functionName string
var startTime string
var endTime string
var duration int32

func NewLogsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logs",
		Short: "logs",
		Long:  "logs",
		Run:   runLogs,
	}
	cmd.Flags().StringVarP(&functionName, "name", "n", "", "Function name.")
	cmd.Flags().StringVarP(&startTime, "start-time", "s", "", "log start time.")
	cmd.Flags().StringVarP(&endTime, "end-time", "e", "", "log end time.")
	cmd.Flags().Int32VarP(&duration, "duration", "d", 0, "count of logs.")

	return cmd
}

func runLogs(cmd *cobra.Command, args []string) {
	if functionName == "" {
		fmt.Printf("function name is empty.")
		return
	}
	// get user info
	user := common.GetUser()
	if logSetId, logTopicId, err := getFunction(user, functionName); err == nil {
		findLog(user, logSetId, logTopicId)
	}
}

func getFunction(user *common.User, functionName string) (logSetId, logTopicId string, err error) {
	req := functionApis.NewGetFunctionRequestWithAllParams(user.Region, functionName)
	resp, err := common.NewFunctionClient(user).GetFunction(req)
	if err != nil || resp.Error.Code != 0 {
		if err != nil {
			fmt.Printf("Get function (%s) error : %s \n", functionName, err.Error())
		} else {
			fmt.Printf("Get function (%s) error : %s \n", functionName, resp.Error.Message)
		}
		return "", "", err
	}
	return resp.Result.Data.LogSetId, resp.Result.Data.LogTopicId, nil
}

type FunctionContent struct {
	RequestId    string `json:"request_id"`
	FunctionName string `json:"function_name"`
	Version      string `json:"version"`
	Content      string `json:"content"`
	Message      string `json:"message"`
}

func findLog(user *common.User, logSetId, logTopicId string) {
	// https://docs.jdcloud.com/cn/log-service/api/search?content=API
	client := common.NewLogClient(user)
	req := logsApis.NewSearchRequest(user.Region, logSetId, logTopicId, "fulltext")

	var err error
	var start *time.Time
	var end *time.Time
	now := time.Now()
	if duration > 0 {
		start = &time.Time{}
		end = &time.Time{}

		*end = now
		*start = end.Add(time.Duration(-duration) * time.Second)
	} else {
		duration = 600
		if startTime != "" {
			start = &time.Time{}
			if *start, err = time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local); err != nil {
				fmt.Printf("start-time(%s),parse err=%s", startTime, err.Error())
				return
			}
		}
		if endTime != "" {
			end = &time.Time{}
			if *end, err = time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local); err != nil {
				fmt.Printf("end-time(%s),parse err=%s", endTime, err.Error())
				return
			}
		}
		if start == nil {
			start = &time.Time{}
			*start = end.Add(time.Duration(-duration) * time.Second)
		}
		if end == nil {
			end = &time.Time{}
			*end = start.Add(time.Duration(duration) * time.Second)
		}
	}
	req.SetStartTime(start.Format("2006-01-02T15:04:05Z0700"))
	req.SetEndTime(end.Format("2006-01-02T15:04:05Z0700"))

	pageSize := 50
	req.SetPageSize(pageSize)
	currentPageNumber := 1
	for {
		req.SetPageNumber(currentPageNumber)
		resp, err := client.Search(req)
		if err != nil {
			fmt.Printf("find log err=%s\n", err.Error())
			return
		}
		if resp.Error.Code != 0 && resp.Error.Code != 200 {
			fmt.Printf("find log err=%s\n", resp.Error.Message)
			return
		}

		for _, val := range resp.Result.Data {
			if data, ok := val.(map[string]interface{}); ok {
				t := int64(data["time"].(float64))
				content := data["content"]
				funContent := &FunctionContent{}
				if err := json.Unmarshal([]byte(content.(string)), funContent); err == nil {
					if funContent.Content != "" {
						fmt.Printf("%s %s %s\n", time.Unix(0, t*1e6).Format(time.RFC3339), funContent.RequestId, funContent.Content)
					} else {
						fmt.Printf("%s %s %s\n", time.Unix(0, t*1e6).Format(time.RFC3339), funContent.RequestId, funContent.Message)
					}
				} else {
					fmt.Printf("unmarshal (%s),err=%s\n", content, err.Error())
				}
			}
		}

		pages := int(resp.Result.Total) / pageSize
		if int(resp.Result.Total)%pageSize > 0 {
			pages += 1
		}
		if currentPageNumber < pages {
			yes := "y"
			fmt.Printf("\n[PageNumber:%d/Total:%d]\n", currentPageNumber, pages)
			fmt.Println("Continue to print or not? Y/N")
			fmt.Scanln(&yes)
			if yes == "y" || yes == "Y" {
				currentPageNumber += 1
			} else {
				return
			}
		} else {
			return
		}
	}
}
