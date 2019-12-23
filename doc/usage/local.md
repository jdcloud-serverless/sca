# sca local

在deploy函数之前，可以在本地调试函数

在本地调试之前，可以手动下载本地调试环境镜像(当前支持python的三个版本, sca也会自动下载)。镜像包含函数的运行环境，与线上的环境是一致的。  
`docker pull sca1/python27`  
`docker pull sca1/python36`  
`docker pull sca1/python37`  

## 用法

`./sca local -t ./template.yaml -e ./event.json function_name`

`template.yaml`文件可以通过`sca init`命令生成模板，针对函数属性对模板进行修改即可。

`event.json`文件保存本次调用的事件信息，必须是json格式

`function_name`模板文件中的函数名称
