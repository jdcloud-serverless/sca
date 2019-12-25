# sca function list

sca function list 命令可以查看云端已存在函数资源的列表

## 用法

`sca function list`

## 示例
```
# sca function list
      +--------------------------------------------+-------------------------------------------------+-----------+-------------------------------------------------------------------------------------------------------------------------------------+----------------------+
<<<<<<< HEAD
      |               FUNCTION NAME                |                   DESCRIPTION                   |  RUNTIME  |                                                              CODE URL                                                                                                                  |     CREATE TIME      |
      +--------------------------------------------+-------------------------------------------------+-----------+-------------------------------------------------------------------------------------------------------------------------------------+----------------------+
      | test-function                              | This is a template of function                  | python3.6 | http://oss-function-hb.s3.cn-north-1.jcloudcs.com/jcloudiaas2%3Atest-function%3ALATEST.zip                                          | 2019-12-24T03:30:54Z |
      |                                            | which name is "test-function"                   |           |                                                                                                                                                                                        |                      |
      | function-sca1224                           | This is a template of function                  | python3.6 | http://oss-function-hb.s3.cn-north-1.jcloudcs.com/jcloudiaas2%3Afunction-sca1224%3ALATEST.zip                                       | 2019-12-24T03:07:09Z |
      |                                            | which name is "test-function"                   |           |                                                                                                                                                                                        |                      |
=======
      |               FUNCTION NAME                |                   DESCRIPTION                   |  RUNTIME  |                                                              CODE URL                                                               |     CREATE TIME      |
      +--------------------------------------------+-------------------------------------------------+-----------+-------------------------------------------------------------------------------------------------------------------------------------+----------------------+
      | test-function                              | This is a template of function                  | python3.6 | http://oss-function-hb.s3.cn-north-1.jcloudcs.com/jcloudiaas2%3Atest-function%3ALATEST.zip                                          | 2019-12-24T03:30:54Z |
      |                                            | which name is "test-function"                   |           |                                                                                                                                     |                      |
      | function-sca1224                           | This is a template of function                  | python3.6 | http://oss-function-hb.s3.cn-north-1.jcloudcs.com/jcloudiaas2%3Afunction-sca1224%3ALATEST.zip                                       | 2019-12-24T03:07:09Z |
      |                                            | which name is "test-function"                   |           |                                                                                                                                     |                      |
>>>>>>> fix usage doc for output format
      +--------------------------------------------+-------------------------------------------------+-----------+-------------------------------------------------------------------------------------------------------------------------------------+----------------------+

```
