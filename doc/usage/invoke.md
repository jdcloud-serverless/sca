# sca invoke

sca invoke 命令可以远程调用云端已经存在的函数资源

## 用法

 参数 | 简写 | 必填 | 描述 | 示例
 ------------ |------------| ------------|------------|------------
 name|-n|是|调用指定函数|-n myfunction
 event|-e|否|指定测试模版文件，若不指定则测试模板则默认传 key:value 字符串	|-e ./event.json


## 示例
```
# sca invoke -n sca-001 -e ./event.json
2019-12-23T18:12:01+08:00  bp095tamssuc3eknva3tktf77a3jjjh1  Start Invoke
2019-12-23T18:12:01+08:00  bp095tamssuc3eknva3tktf77a3jjjh1  {u'key_001': u'value_001'}
2019-12-23T18:12:01+08:00  bp095tamssuc3eknva3tktf77a3jjjh1  End Invoke
2019-12-23T18:12:01+08:00  bp095tamssuc3eknva3tktf77a3jjjh1  Report Invoke bp095tamssuc3eknva3tktf77a3jjjh1,Duration :8.21ms  BilledDuration: 100ms  Memory Size: 128 MB

Invoke Return : hello world
RequestId: bp095tamssuc3eknva3tktf77a3jjjh1              Billed Duration: 100 ms                Memory Size: 128 MB             Max Memory Used : 0.12 MB

```