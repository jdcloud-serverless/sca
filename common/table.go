package common

import (
	"os"
	"strconv"

	"github.com/jcloud-api-gateway/jcloud-sdk-go/services/function/models"

	"github.com/olekukonko/tablewriter"
)

func TableFunctionModel(function models.Function) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Function Name", "Description", "Version", "Runtime", "Timeout", "Memory Size", "Handler", "Code Url", "Create Time"})
	row := []string{}
	row = append(row, function.Name)
	row = append(row, function.Description)
	row = append(row, function.Version)
	row = append(row, function.RunTime)
	row = append(row, strconv.Itoa(function.OverTime)+" s")
	row = append(row, strconv.Itoa(function.Memory)+" MB")
	row = append(row, function.Entrance)
	row = append(row, function.DownloadUrl)
	row = append(row, function.CreateTime)
	table.Append(row)
	table.Render()
}
