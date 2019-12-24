package template

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	DefaultROSTemplateFormatVersion = "2019-12-25"
	DefaultTransform                = "JDCloud::Serverless-2019-12-25"
	DefaultFunctionType             = "JDCloud::Serverless::Function"

	DefaultFunctionName = "test-function"
	DefaultProjectName  = "testproject"

	RUNTIME_Python2_7 = "python2.7"
	RUNTIME_Python3_6 = "python3.6"
	RUNTIME_Python3_7 = "python3.7"
)

type Template struct {
	ROSTemplateFormatVersion string                       `yaml:"ROSTemplateFormatVersion"`
	Transform                string                       `yaml:"Transform"`
	Resources                map[string]*FunctionTemplate `yaml:"Resources"`
}

type FunctionTemplate struct {
	Type               string             `yaml:"Type"`
	FunctionProperties FunctionProperties `yaml:"Properties"`
}

type FunctionProperties struct {
	Handler     string            `yaml:"Handler"`
	Timeout     int               `yaml:"Timeout"`
	MemorySize  int               `yaml:"MemorySize"`
	Runtime     string            `yaml:"Runtime"`
	Description string            `yaml:"Description"`
	CodeUri     string            `yaml:"CodeUri"`
	Env         map[string]string `yaml:"Env"`

	// TODO
	Role      string    `yaml:"Role"`
	Policies  string    `yaml:"Policies"`
	VPCConfig VPCConfig `yaml:"VPCConfig"`
	LogConfig LogConfig `yaml:"LogConfig"`
	//Events    []Event   `yaml:"Events"`
}

type VPCConfig struct {
	Vpc    string `yaml:"Vpc"`
	Subnet string `yaml:"Subnet"`
}

type LogConfig struct {
	LogSet   string `yaml:"LogSet"`
	LogTopic string `yaml:"LogTopic"`
}

type Event struct {
}

func LoadTemplate(filename string) (*Template, error) {
	conf := new(Template)
	err := load(filename, conf)
	if err != nil {
		fmt.Printf("Load template error = %s\nPlease input correct template path\n", err.Error())
		return nil, err
	}
	return conf, nil
}

func load(filename string, in interface{}) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, in)
	if err != nil {
		return err
	}

	return nil
}

func Marshal(tmpl *Template) (out []byte, err error) {
	return yaml.Marshal(tmpl)
}
