# SCA (Serverless Cloud Application)

## 简介

[SCA CLI](https://github.com/jdcloud-serverless/sca)是京东云无服务器云应用（Serverless Cloud Application，SCA）命令行工具。对京东云serverless应用服务提供更加便捷的本地管理功能，包括本地函数管理：本地测试、打包、部署、云端测试等。          
sca cli 通过一个函数模板配置文件template.yaml，完成函数资源的描述，并基于配置文件实现本地代码及配置部署到云端的过程。

## 运行环境

sca cli支持Linux、Mac运行；    
sca cli基于go开发完成，您只需下载安装包，即可使用。    

## 开始使用

### 安装 sca cli（Linux）

执行以下命令一步完成下载安装：           
` curl -O https://raw.githubusercontent.com/jdcloud-serverless/sca/master/hack/install.sh && chmod +777 install.sh && sh install.sh && source ~/.bashrc`     

此外，您可以[下载安装包至本地](https://github.com/jdcloud-serverless/sca/releases)后，执行`chmod +x sca`命令给予可执行权限后运行。

### 安装 sca cli（Mac）

执行以下命令一步完成下载安装：    

` curl -O https://raw.githubusercontent.com/jdcloud-serverless/sca/master/hack/install-mac.sh && chmod +x install-mac.sh && sh install-mac.sh && source ~/.bashrc  `


此外，您可以[下载安装包至本地](https://github.com/jdcloud-serverless/sca/releases)后，执行`chmod +x sca`命令给予可执行权限后运行。  



### 查询sca版本
` sca version `      
`JD Serverless Cloud Application Version: 0.0.1`

### 配置账号信息  
sca安装完成后，进行[初始化配置](https://github.com/jdcloud-serverless/sca/blob/master/doc/usage/config.md)，将JDCloud的账号信息同步至sca中。 

```
#sca config
[>] JDCould accountid = 11111(your acount id)
[>] JDCould region = cn-north-1
[>] JDCould access-key = 0123abcd(your ak)
[>] JDCould secret-key = abcd0123(your sk)

```       



### 安装Docker(执行local命令必须)

sca cli 支持使用 Docker 容器管理工具启动和使用容器，作为在本地运行函数代码的环境。sca cli的 local 命令将会使用 Docker 的管理接口实现相关交互。     

- 如果您需要使用本地调试、运行能力，请确保 Docker 已正确安装。                          
- 如果您当前不需要使用 Docker 或者计划稍后再安装 Docker 时，可跳过此步骤。                       

#### 在 Mac 上安装 Docker
| 版本                                 | 下载地址                                                     | 安装方式                                                     |
| ------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| Apple Mac OS Sierra 10.12 及以上版本 | [docker-ce-desktop](https://hub.docker.com/editions/community/docker-ce-desktop-mac) | 下载 docker.dmg 安装包，双击安装文件，启动安装。             |
| Apple Mac OS Sierra 10.12 以下版本   | [Docker Toolbox](https://docs.docker.com/toolbox/overview/)  | 根据 Toolbox 提供的[ macOS 安装指导](https://docs.docker.com/toolbox/toolbox_install_mac/)，双击 Toolbox 安装工具，安装 Toolbox。完成安装后，双击 Launchpad 中新增的 Docker Quickstart Terminal 图标，启动 Docker 。 |

#### 在 Linux 上安装 Docker


| 版本            | 下载地址                                                     | 安装方式                                                     |
| --------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| CentOS 操作系统 | [CentOS](https://docs.docker.com/install/linux/docker-ce/centos/) | 执行 sudo yum install docker-ce docker-ce-cli containerd.io 命令，安装 Docker 。 |
| Debian 操作系统 | [Debian](https://docs.docker.com/install/linux/docker-ce/debian/) | 执行 sudo apt-get install docker-ce docker-ce-cli containerd.io 命令，安装 Docker。 |
| Fedora 操作系统 | [Fedora](https://docs.docker.com/install/linux/docker-ce/fedora/) | 执行 sudo dnf install docker-ce docker-ce-cli containerd.io 命令，安装 Docker。 |
| Ubuntu 操作系统 | [Ubuntu](https://docs.docker.com/install/linux/docker-ce/ubuntu/) | 执行 sudo apt-get install docker-ce docker-ce-cli containerd.io 命令，安装 Docker。 |
| 二进制包        | [二进制包](https://docs.docker.com/install/linux/docker-ce/binaries/) | 解压并运行二进制包，即可完成 Docker 的下载安装和启动。       |



### 初始化项目       
通过 [初始化项目](https://github.com/jdcloud-serverless/sca/blob/master/doc/usage/init.md) ，用户可快速创建一个简单的模板，包括代码文件、配置文件，基于模板可进一步进行配置及开发后，可直接打包部署云端。     

示例：在当前目录初始化创建testproject（默认）工程，默认函数名称：test-function，运行时：python3.6，您也可以通过配置参数创建：

`  sca init   `  

创建工程目录如下

```
[root@A04-R08-I139-110-7T9CT92 testproject]# tree
.
├── README.md
├── template.yaml
└── test-function
    └── index.py

1 directory, 3 files

```


### 配置文件template.yaml
初始化命令会在工程中创建一个template.yaml模板，sca cli 通过此函数模板配置文件，完成函数资源的描述，并基于配置文件实现本地代码及配置部署到云端。格式如下：

```
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

### 验证配置文件
 [验证template.yaml文件](https://github.com/jdcloud-serverless/sca/blob/master/doc/usage/validate.md)    
 
```
# sca validate -t template.yaml
validate success.
```

### 本地测试
通过 [本地调试函数](https://github.com/jdcloud-serverless/sca/blob/master/doc/usage/local.md) ，在部署前，用户可在本地的模拟环境中运行代码，发送模拟测试事件，验证函数执行，获取运行信息及日志。在运行本地调试前，需确保本地环境中已经安装并启动 Docker。             

示例：测试本地默认当前目录下template.yaml文件中的test-function函数，event测试事件则默认 key:value 字符串：

```
#  sca local -n test-function

```  

```
python36: Pulling from jdccloudserverless/sca
Digest: sha256:6c40080bf12f45881a1f92e865eb52895a4a694ad4b12f620f25d8e95d52c6bd
Status: Image is up to date for jdccloudserverless/sca:python36
{
	"code": 0,
	"return": "hello world",
	"stdout": "{}\n",
	"stderr": "",
	"memory_used": "0.00m",
	"time_used": "6.889629ms"
}

```



### 打包部署
根据指定的函数模板配置文件，将配置文件中的指定代码包、函数配置等信息， [打包部署到云端](https://github.com/jdcloud-serverless/sca/blob/master/doc/usage/deploy.md) 。 

示例：将当前目录template.yaml配置中的函数部署至云端，默认覆盖云端重名函数：
```
sca deploy 
```

```
Function (test-function) not exists , now beginning to create function
+---------------+--------------------------------+---------+-----------+---------+-------------+---------------+--------------------------------------------------------------------------------------------+----------------------+
| FUNCTION NAME |          DESCRIPTION           | VERSION |  RUNTIME  | TIMEOUT | MEMORY SIZE |    HANDLER    |                                          CODE URL                                          |     CREATE TIME      |
+---------------+--------------------------------+---------+-----------+---------+-------------+---------------+--------------------------------------------------------------------------------------------+----------------------+
| test-function | This is a template of function | LATEST  | python3.6 | 300 s   | 128 MB      | index.handler | http://oss-function-hb.s3.cn-north-1.jcloudcs.com/xxxxxx%3Atest-function%3ALATEST.zip | 2019-12-30T03:10:32Z |
|               | which name is "test-function"  |         |           |         |             |               |                                                                                            |                      |
+---------------+--------------------------------+---------+-----------+---------+-------------+---------------+--------------------------------------------------------------------------------------------+----------------------+
Deploy function (test-function) success .
```



### 函数管理
通过函数管理，您可以[查看云端函数列表](https://github.com/jdcloud-serverless/sca/blob/master/doc/usage/function_list.md)、[查询函数配置](https://github.com/jdcloud-serverless/sca/blob/master/doc/usage/function_info.md)，并可以[删除函数](https://github.com/jdcloud-serverless/sca/blob/master/doc/usage/function_delete.md)。               
`sca function list`  查看云端已存在函数资源的列表。   
`sca function info`  查看已部署云端函数配置。             
`sca function del`   删除已部署云端函数。     

示例：查询云端test-function函数配置详情：
``` 
sca function info -n test-function   
```
```
+---------------+--------------------------------+---------+-----------+---------+-------------+---------------+--------------------------------------------------------------------------------------------+----------------------+
| FUNCTION NAME |          DESCRIPTION           | VERSION |  RUNTIME  | TIMEOUT | MEMORY SIZE |    HANDLER    |                                          CODE URL                                          |     CREATE TIME      |
+---------------+--------------------------------+---------+-----------+---------+-------------+---------------+--------------------------------------------------------------------------------------------+----------------------+
| test-function | This is a template of function | LATEST  | python3.6 | 300 s   | 128 MB      | index.handler | http://oss-function-hb.s3.cn-north-1.jcloudcs.com/xxxxxx%3Atest-function%3ALATEST.zip | 2019-12-30T03:10:32Z |
|               | which name is "test-function"  |         |           |         |             |               |                                                                                            |                      |
+---------------+--------------------------------+---------+-----------+---------+-------------+---------------+--------------------------------------------------------------------------------------------+----------------------+

```



### 云端调用函数
通过invoke命令用户可于本地[调用云端函数](https://github.com/jdcloud-serverless/sca/blob/master/doc/usage/invoke.md)，进行测试验证。

示例：本地测试云端test-function函数：
```
sca invoke -n test-function

```

```
2019-12-30T11:14:15+08:00  bp4no1qq6sh2v13hpoed8vdpfqr84n05  Start Invoke   
2019-12-30T11:14:15+08:00  bp4no1qq6sh2v13hpoed8vdpfqr84n05  {'key': 'value'}  
2019-12-30T11:14:15+08:00  bp4no1qq6sh2v13hpoed8vdpfqr84n05  End Invoke   
2019-12-30T11:14:15+08:00  bp4no1qq6sh2v13hpoed8vdpfqr84n05  Report Invoke bp4no1qq6sh2v13hpoed8vdpfqr84n05,Duration :7.54ms  BilledDuration: 100ms  Memory Size: 128 MB  

Invoke Return : hello world
RequestId: bp4no1qq6sh2v13hpoed8vdpfqr84n05 		 Billed Duration: 100 ms 		Memory Size: 128 MB 		Max Memory Used : 0.00 MB


```


### 云端日志
通过[查询云端日志命令](https://github.com/jdcloud-serverless/sca/blob/master/doc/usage/logs.md)，您可以查询指定云端函数某时段内的执行日志。

说明：查询云端日志，请先为函数配置日志集及日志主题，在template.yaml文件中配置LogSetID和LogTopicID，方可通过日志服务查询函数执行日志：
```    
      LogConfig:
        LogSet: "LogSetID"
        LogTopic: "LogTopicID"
````

示例：查询test-function函数最近600000秒日志：
```
# sca logs -n test-function -d 600000
2019-12-19T10:35:05+08:00 boue3nfsqrshctda7hp792adjrap4r6r Report Invoke boue3nfsqrshctda7hp792adjrap4r6r,Duration :7.54ms  BilledDuration: 100ms  Memory Size: 128 MB
2019-12-19T10:35:05+08:00 boue3nfsqrshctda7hp792adjrap4r6r End Invoke
2019-12-19T10:35:05+08:00 boue3nfsqrshctda7hp792adjrap4r6r {}
2019-12-19T10:35:05+08:00 boue3nfsqrshctda7hp792adjrap4r6r Start Invoke

......


```

