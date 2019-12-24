package initialize

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jdcloud-serverless/sca/common"

	"github.com/spf13/cobra"
)

var runtime string
var output string
var projectName string

const indexPthon = `def handler(event,context):
    print(event)
    return "hello world"
`

const tmplContent = `Resources:
  default:
      Type: TencentCloud::Serverless::Function
      Properties:
        CodeUri: ./
        Type: Event
        Description: This is a template function
        Handler: index.handler
        MemorySize: 128
        Runtime: Python3.6
        Timeout: 3
        Environment:
          Variables:
            ENV_FIRST: env1
            ENV_SECOND: env2
        VpcConfig:
           VpcId: ' '
           SubnetId: ' '
       LogConfig:
            `

func NewInitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "init a function template",
		Long:  "init a function template",
		Run:   initFun,
	}
	cmd.Flags().StringVarP(&runtime, "runtime", "r", "", "Runtime of this funtion.Include python3.6,python3.7,python2.7")
	cmd.Flags().StringVarP(&output, "output-dir", "o", "", "The path where will output the initialized app into.")
	cmd.Flags().StringVarP(&projectName, "name", "n", "", "project name.")
	return cmd
}

func initFun(cmd *cobra.Command, args []string) {
	if runtime == "" {
		runtime = common.RUNTIME_Python3_6
	} else {
		switch runtime {
		case common.RUNTIME_Python2_7, common.RUNTIME_Python3_6, common.RUNTIME_Python3_7:
		default:
			fmt.Printf("%s runtime is not support.\n", runtime)
			return
		}
	}
	if output == "" {
		output, _ = os.Getwd()
	}else {
		if !filepath.IsAbs(output){
			currentPath,_ :=os.Getwd()
			output = fmt.Sprintf("%s/%s",currentPath,output)
		}
	}
	if projectName == "" {
		projectName = common.DefaultProjectName
	}

	funPath := fmt.Sprintf("%s/%s", output, projectName)
	if err := os.MkdirAll(funPath, os.ModePerm); err != nil {
		fmt.Printf("create path err=%s\n", err.Error())
	}

	readmeFile, err := os.Create(filepath.Join(funPath, "README.md"))
	if err != nil {
		fmt.Printf("create README.md err=%s\n", err.Error())
	}
	defer readmeFile.Close()
	readmeFile.WriteString("")

	if writeRootFile(funPath) != nil {
		return
	}

	templateFile, err := os.Create(filepath.Join(funPath, "template.yaml"))
	if err != nil {
		fmt.Printf("create template.yaml err=%s\n", err.Error())
	}
	defer templateFile.Close()
	tmpl := common.Template{
		ROSTemplateFormatVersion: common.DefaultROSTemplateFormatVersion,
		Transform:                common.DefaultTransform,
		Resources:                map[string]*common.FunctionTemplate{},
	}
	tmpl.Resources[common.DefaultFunctionName] = &common.FunctionTemplate{
		Type: common.DefaultFunctionType,
		FunctionProperties: common.FunctionProperties{
			Handler:     "index.handler",
			Timeout:     300,
			MemorySize:  128,
			Runtime:     runtime,
			Description: fmt.Sprintf("This is a template of function which name is \"%s\"", common.DefaultFunctionName),
			CodeUri:     "./",
		},
	}
	writeTemplate(tmpl.Resources[projectName])
	if o, err := common.Marshal(&tmpl); err != nil {
		fmt.Printf("marsharl template.yaml err=%s\n", err.Error())
	} else {
		templateFile.Write(o)
	}
}

func writeRootFile(funPath string) error {
	switch runtime {
	case common.RUNTIME_Python2_7, common.RUNTIME_Python3_6, common.RUNTIME_Python3_7:
		return writeRootFile_python(funPath, runtime)
	default:
		return fmt.Errorf("%s runtime is not support.", runtime)
	}
	return nil
}

func writeRootFile_python(funPath, runtime string) error {
	rootFile, err := os.Create(filepath.Join(funPath, "index.py"))
	if err != nil {
		fmt.Printf("create index.py err=%s\n", err.Error())
	}
	defer rootFile.Close()
	rootFile.WriteString(indexPthon)
	return nil
}

func writeTemplate(tmpl *common.FunctionTemplate) {
	switch runtime {
	case common.RUNTIME_Python2_7, common.RUNTIME_Python3_6, common.RUNTIME_Python3_7:
		writeTemplate_python(tmpl)
	default:

	}
}

func writeTemplate_python(tmpl *common.FunctionTemplate) {

}
