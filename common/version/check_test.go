package version

import (
	"fmt"
	"testing"
)

func TestStartCheckVersion(t *testing.T) {
	code, msg := StartCheckVersion()
	fmt.Printf("code: %v\n", code)
	fmt.Printf("message: %v\n", msg)
}
