# SCA

## 什么是SCA

SCA （Serverless Cloud Application）是京东云函数计算 （Function）产品的命令行工具。
通过 SCA 命令行工具，您可以方便的实现函数打包、部署、本地调试，也可以方便的生成函数计算的项目并基于 demo 项目进一步的开发。

SCA 通过一个资源配置文件（template.yaml），协助您进行开发、构建、部署操作。

## 安装

`curl -O https://raw.githubusercontent.com/jdcloud-serverless/sca/master/hack/install.sh && chmod +777 install.sh && sh install.sh && source ~/.bashrc`


## 开始使用

通过 sca 命令行工具，你可以：

[配置](https://github.com/jdcloud-serverless/sca/tree/master/doc/usage/config.md)用户信息；

快速[初始化](https://github.com/jdcloud-serverless/sca/tree/master/doc/usage/init.md)函数计算项目；

[验证](https://github.com/jdcloud-serverless/sca/tree/master/doc/usage/validate.md)资源配置文件；

在[本地](https://github.com/jdcloud-serverless/sca/tree/master/doc/usage/local.md)开发及测试你的函数计算代码；

[部署](https://github.com/jdcloud-serverless/sca/tree/master/doc/usage/deploy.md)函数代码，创建函数及更新函数配置；

查询函数[信息](https://github.com/jdcloud-serverless/sca/tree/master/doc/usage/function_info.md)，获取函数[列表](https://github.com/jdcloud-serverless/sca/tree/master/doc/usage/function_list.md)，[删除](https://github.com/jdcloud-serverless/sca/tree/master/doc/usage/function_delete.md)指定函数；

[调用](https://github.com/jdcloud-serverless/sca/tree/master/doc/usage/invoke.md)云端函数，获取执行结果；

查看云端函数的执行[日志](https://github.com/jdcloud-serverless/sca/tree/master/doc/usage/logs.md)。


## 反馈

如您在使用中遇到问题，可以在 issue 反馈问题。