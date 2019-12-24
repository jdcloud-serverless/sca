# sca init

快速初始化

## 用法

 参数 | 简写 | 必填 | 描述 | 示例
 ------------|------------| ------------|------------|------------
 name|-n|是|生成的项目名称。如果不填写，默认值为 testproject。|-n testproject
 runtime|-r|否|生成的项目运行环境，可选值为python2.7、python3.6、3.7。默认值为 python3.6。	|-r python2.7
output-dir|-o|否|指定项目生成的目录，默认为当前目录。|-o /root/sca


## 备注
## 示例
```
# sca init -n testproject -r python3.6 -o mysca

# ls -l mysca/testproject/
总用量 2
-rwxrwx--- 1 root root  70 12月 23 17:08 index.py
-rwxrwx--- 1 root root   0 12月 23 17:08 README.md
-rwxrwx--- 1 root root 530 12月 23 17:08 template.yaml

# cat mysca/testproject/index.py
def handler(event,context):
    print(event)
    return "hello world"

# cat mysca/testproject/template.yaml
ROSTemplateFormatVersion: "2019-12-25"
Transform: JDCloud::Serverless-2019-12-25
Resources:
  test-function:
    Type: JDCloud::Serverless::Function
    Properties:
      Handler: index.handler
      Timeout: 300
      MemorySize: 128
      Runtime: python3.6
      Description: This is a template of function which name is "test-function"
      CodeUri: ./
      Env: {}
      Role: ""
      Policies: ""
      VPCConfig:
        Vpc: ""
        Subnet: ""
      LogConfig:
        LogSet: ""
        LogTopic: ""
```