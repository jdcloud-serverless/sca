# sca deploy

sca deploy 命令可以将template.yaml文件定义的函数资源部署到云端。

## 用法

`.sca/deploy -t ./template.yaml` 或  `.sca/deploy --template ./template.yaml`

`template.yaml`文件可以通过`sca init`命令生成模板，针对函数属性对模板进行修改即可。

如果`template.yaml`文件中对应的函数资源在云端已经存在，`sca deploy` 命令会对云端函数进行更新，如果函数资源在云端不存在，`sca deploy`命令会在云端创建对应函数资源。

## 示例
```
# sca deploy -t mysca/testproject/template.yaml
Function (test-function) not exists , now beginning to create function
+---------------+--------------------------------+---------+-----------+---------+-------------+---------------+--------------------------------------------------------------------------------------------+----------------------+
| FUNCTION NAME |          DESCRIPTION           | VERSION |  RUNTIME  | TIMEOUT | MEMORY SIZE |    HANDLER    |                                          CODE URL                                          |     CREATE TIME      |
+---------------+--------------------------------+---------+-----------+---------+-------------+---------------+--------------------------------------------------------------------------------------------+----------------------+
| test-function | This is a template of function | LATEST  | python3.6 | 3 s     | 128 MB      | index.handler | http://oss-function-hb.s3.cn-north-1.jcloudcs.com/jcloudiaas2%3Atest-function%3ALATEST.zip | 2019-12-23T09:45:31Z |
|               | which name is "test-function"  |         |           |         |             |               |                                                                                            |                      |
+---------------+--------------------------------+---------+-----------+---------+-------------+---------------+--------------------------------------------------------------------------------------------+----------------------+
Deploy function (test-function) success .

```