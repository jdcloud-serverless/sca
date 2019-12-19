package local

import (
	"fmt"
	"testing"

	"github.com/jdcloud-serverless/sca/common"
)

func TestExecute(t *testing.T) {
	envs := make(map[string]string)
	envs["key1"] = "value1"

	properties := common.FunctionProperties{
		Name:        "test-function",
		Handler:     "index.handler",
		Timeout:     100,
		MemorySize:  512,
		Runtime:     "python27",
		Description: "a test function",
		CodeUri:     "/root/code",
		Env:         envs,
	}

	eventStr := "{\"key\":\"value\"}"
	event := []byte(eventStr)

	resp := Execute(properties, event)
	fmt.Printf("resp: %v\n", resp)
}
