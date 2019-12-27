package version

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/jdcloud-serverless/sca/common/client"
)

const (
	FunctionApiVersionUrl = "http://w6mzlznt62zn.cn-north-1.jdcloud-api.net/sca/version"

	CurrentVersion    = "1.0"
	UpgradeVersionMsg = "Newest Function Api version: %s, Please download new sca!!!"

	SuccessCode        = 0
	UpgradeVersionCode = 1
	IgnoreErrorCode    = 2

	HttpWaitTime = 1
)

func CheckFunctionApiVersion() (int, string) {
	httpClient := client.NewHttpClient()

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-time.After(time.Second * time.Duration(HttpWaitTime)):
			cancel()
		}
	}()
	resp, err := httpClient.Forward(FunctionApiVersionUrl, http.MethodGet, nil, nil, ctx)
	if err != nil {
		return IgnoreErrorCode, ""
	}
	if resp.StatusCode != http.StatusOK {
		return IgnoreErrorCode, ""
	}
	version, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return IgnoreErrorCode, ""
	}

	newVersion := strings.TrimSpace(string(version))
	if newVersion == CurrentVersion {
		return SuccessCode, ""
	}

	return UpgradeVersionCode, fmt.Sprintf(UpgradeVersionMsg, newVersion)
}
