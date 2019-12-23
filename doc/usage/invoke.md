# sca invoke

sca invoke 命令可以远程调用云端已经存在的函数资源

## 用法

`sca invoke -e event.json -n function_name`

`-e` 参数指定json格式的触发事件的路径

`-n` 参数指定待执行的函数名称

## 示例
```
sca invoke -n sca-001 -e ./event.json
2019-12-23T18:12:01+08:00  bp095tamssuc3eknva3tktf77a3jjjh1  Start Invoke
2019-12-23T18:12:01+08:00  bp095tamssuc3eknva3tktf77a3jjjh1  {u'key_001': u'value_001'}
2019-12-23T18:12:01+08:00  bp095tamssuc3eknva3tktf77a3jjjh1  End Invoke
2019-12-23T18:12:01+08:00  bp095tamssuc3eknva3tktf77a3jjjh1  Report Invoke bp095tamssuc3eknva3tktf77a3jjjh1,Duration :8.21ms  BilledDuration: 100ms  Memory Size: 128 MB

Invoke Return : hello world
RequestId: bp095tamssuc3eknva3tktf77a3jjjh1              Billed Duration: 100 ms                Memory Size: 128 MB             Max Memory Used : 0.12 MB

```