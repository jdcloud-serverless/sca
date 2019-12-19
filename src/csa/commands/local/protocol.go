package local

const (
	HeaderContentType      = "Content-Type"
	TypeJson               = "application/json"
	HeaderRequestId        = "x-fc-request-id"
	HeaderFunctionId       = "x-fc-function-id"
	HeaderFunctionName     = "x-fc-function-name"
	HeaderFunctionVersion  = "x-fc-function-version"
	HeaderFunctionHandler  = "x-fc-function-handler"
	HeaderFunctionMemory   = "x-fc-function-memory"
	HeaderFunctionTimeout  = "x-fc-function-timeout"
	HeaderFunctionMD5      = "x-fc-function-md5"
	HeaderFunctionCodePath = "x-fc-function-codepath"
	HeaderFunctionEnv      = "x-fc-function-env"

	InternalSuccessCode   = 0
	InternalErrorCode     = -1
	InvokeFunctionAskCode = 1004
	UserFuncExecuteError  = 2000 //user function throw exception
)

type WsgiResponse struct {
	StdoutReturn   string  `json:"StdoutReturn"`
	StderrReturn   string  `json:"StderrReturn"`
	FunctionReturn string  `json:"FunctionReturn"`
	MemUsage       float64 `json:"MemUsage"`
}
