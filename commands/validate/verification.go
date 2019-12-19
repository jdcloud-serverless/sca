package validate

import (
	"os"
	"regexp"
	"strings"

	"github.com/jdcloud-serverless/sca/common"
)

const (
	MaxTime               = 300
	MaxFunctionNameLength = 128
)

func RuntimeCheck(runtime string) bool {
	switch runtime {
	case common.RUNTIME_Python2_7, common.RUNTIME_Python3_6, common.RUNTIME_Python3_7:
		return true
	default:
		return false
	}
}

func MemoryCheck(memory int) bool {
	if memory < 128 || memory > 3072 {
		return false
	}

	if memory%64 != 0 {
		return false
	}
	return true
}

func OvertimeCheck(overtime int) bool {
	if overtime < 1 || overtime > MaxTime {
		return false
	}
	return true
}

func EnvCheck(envs map[string]string) bool {
	for k, v := range envs {
		if k == "" || v == "" {
			return false
		}
	}
	return true
}

func FunctionNameCheck(name string) bool {
	nameLength := len(name)
	if nameLength == 0 || nameLength > MaxFunctionNameLength {
		return false
	}
	pattern := "^[a-zA-Z0-9_-]+$"
	if ok, _ := regexp.MatchString(pattern, name); !ok {
		return false
	}
	return true
}

func HandlerCheck(runtime, handler string) bool {
	if handler == "" {
		return false
	}
	h := strings.Split(handler, ".")
	if len(h) != 2 {
		return false
	}
	if h[0] == "" || h[1] == "" {
		return false
	}
	return true
}

func CodeUriCheck(codeUri string) bool {
	if codeUri == "" {
		return false
	}
	f, err := os.Open(codeUri)
	if err != nil {
		return false
	}
	f.Close()
	return true
}
