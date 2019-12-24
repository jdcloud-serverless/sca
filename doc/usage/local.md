# sca local
在deploy函数之前，可以在本地调试函数,需要安装并启动docker

## 镜像
可以手动下载本地调试环境镜像(当前支持python的三个版本； 如果没有下载，sca也会自动下载)。镜像包含函数的运行环境，与线上的环境是一致的  
`docker pull sca1/python27`  
`docker pull sca1/python36`  
`docker pull sca1/python37` 

镜像存储在https://hub.docker.com，如果无法下载，可以自行生成本地镜像, 如：  
`cd local/image/python27 && ./gen_image.sh`

## 用法

 参数 | 简写 | 必填 | 描述 | 示例
 ------------ |------------| ------------|------------|------------
 name|-n|是|指定函数名称（template.yaml中的）|-n myfunction
 event|-e|否|指定测试模版文件，若不指定则测试模板则默认传 key:value 字符串	|-e ./event.json
 template|-t|否|项目描述配置文件的路径或文件名，默认为 template.yaml|-n ./template.yaml
 skip-pull-image|-n|否|不下载镜像文件|--skip-pull-image

`sca local --skip-pull-image -t ./template.yaml -e ./event.json -n function_name`

`template.yaml`文件可以通过`sca init`命令生成模板，针对函数属性对模板进行修改即可。

`event.json`文件保存本次调用的事件信息，必须是json格式

`function_name`模板文件中的函数名称

## 示例
```
[root@localhost sca]# sca local --skip-pull-image -t ./helloworld/template.yaml -e ./event.json -n handler
skip pull sca1/python27:latest
{
	"code": 0,
	"return": "hello world",
	"stdout": "{u'k': u'v'}\n",
	"stderr": "",
	"memory_used": "0.11m",
	"time_used": "804.942067ms"
}
```
