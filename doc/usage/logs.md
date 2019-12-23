# sca logs

查看云端函数产生的历史或实时日志

## 用法

 参数 | 简写 | 必填 | 描述 | 示例
 ------------ |------------| ------------|------------|------------
 name|-n|是|获取指定函数的日志|-n myfunction
 duration|-d|否|获取最近 x 秒 的日志	|-d 600
 start-time|-s|否|获取指定开始时间之后的日志，若无end-time，默认10min|-s "2019-7-12 00:00:00"
 end-time|-e|否|获取指定开始时间之前的日志，若无start-time，默认10min	|-e "2019-7-12 00:10:00"

## 备注

* 不输入 --duration 参数，默认获取最近10min的日志。

* 每屏显示50行，超过50行，输入Y，查看下一页日志。

* 时间范围间隔最长1天。

* 可以查询30天内的日志。

* 函数配置了日志采集，才能通过该接口查询日志。

## 示例
```
# sca logs -n sca-001 -d 600000
2019-12-19T10:35:05+08:00 boue3nfsqrshctda7hp792adjrap4r6r Report Invoke boue3nfsqrshctda7hp792adjrap4r6r,Duration :7.54ms  BilledDuration: 100ms  Memory Size: 128 MB
2019-12-19T10:35:05+08:00 boue3nfsqrshctda7hp792adjrap4r6r End Invoke
2019-12-19T10:35:05+08:00 boue3nfsqrshctda7hp792adjrap4r6r {}
2019-12-19T10:35:05+08:00 boue3nfsqrshctda7hp792adjrap4r6r Start Invoke
2019-12-19T10:17:35+08:00 boudqumbcw08a43ri4aw86340n8a85ge Report Invoke boudqumbcw08a43ri4aw86340n8a85ge,Duration :6.03ms  BilledDuration: 100ms  Memory Size: 128 MB
2019-12-19T10:17:35+08:00 boudqumbcw08a43ri4aw86340n8a85ge End Invoke
2019-12-19T10:17:35+08:00 boudqumbcw08a43ri4aw86340n8a85ge {}
2019-12-19T10:17:35+08:00 boudqumbcw08a43ri4aw86340n8a85ge Start Invoke
2019-12-19T10:17:35+08:00 boudrskojdn2sidrihc59obk9dchnf46 Report Invoke boudrskojdn2sidrihc59obk9dchnf46,Duration :7.50ms  BilledDuration: 100ms  Memory Size: 128 MB
2019-12-19T10:17:35+08:00 boudrskojdn2sidrihc59obk9dchnf46 End Invoke
2019-12-19T10:17:35+08:00 boudrskojdn2sidrihc59obk9dchnf46 {}
2019-12-19T10:17:35+08:00 boudrskojdn2sidrihc59obk9dchnf46 Start Invoke
2019-12-18T17:31:32+08:00 botw3dovepe5iek2wap4ue5a6kverf5t Report Invoke botw3dovepe5iek2wap4ue5a6kverf5t,Duration :8.31ms  BilledDuration: 100ms  Memory Size: 128 MB
2019-12-18T17:31:32+08:00 botw3dovepe5iek2wap4ue5a6kverf5t End Invoke
2019-12-18T17:31:32+08:00 botw3dovepe5iek2wap4ue5a6kverf5t {u'base64OwnerPin': u'NTk0MDM1MjYzMDE5', u'resources': [], u'detail': {u'requestContext': {u'sourceIp': u'10.0.2.14', u'requestId': u'c6af9ac6-7b61-11e6-9a41-93e8deadbeef', u'identity': {u'user': u'', u'accountId': u'', u'authType': u'', u'apiKey': u''}, u'stage': u'test', u'apiId': u'testsvc'}, u'body': u'string of request payload', u'headers': {u'header': u'headerValue'}, u'pathParameters': {u'pathParam': u'pathValue'}, u'queryParameters': {u'queryParam': u'queryValue'}, u'path': u'api request path', u'httpMethod': u'GET/POST/DELETE/PUT/PATCH'}, u'source': u'apigateway', u'version': u'0', u'id': u'6a7e8feb-b491-4cf7-a9f1-bf3703467718', u'time': u'2006-01-02T15:04:05.999999999Z', u'detailType': u'ApiGatewayReceived', u'region': u'cn-north-1'}
2019-12-18T17:31:32+08:00 botw3dovepe5iek2wap4ue5a6kverf5t Start Invoke


# sca logs -n sca-001 -s "2019-12-19 10:00:00" -e "2019-12-19 11:00:00"
2019-12-19T10:35:05+08:00 boue3nfsqrshctda7hp792adjrap4r6r Report Invoke boue3nfsqrshctda7hp792adjrap4r6r,Duration :7.54ms  BilledDuration: 100ms  Memory Size: 128 MB
2019-12-19T10:35:05+08:00 boue3nfsqrshctda7hp792adjrap4r6r End Invoke
2019-12-19T10:35:05+08:00 boue3nfsqrshctda7hp792adjrap4r6r {}
2019-12-19T10:35:05+08:00 boue3nfsqrshctda7hp792adjrap4r6r Start Invoke
2019-12-19T10:17:35+08:00 boudqumbcw08a43ri4aw86340n8a85ge Report Invoke boudqumbcw08a43ri4aw86340n8a85ge,Duration :6.03ms  BilledDuration: 100ms  Memory Size: 128 MB
2019-12-19T10:17:35+08:00 boudqumbcw08a43ri4aw86340n8a85ge End Invoke
2019-12-19T10:17:35+08:00 boudqumbcw08a43ri4aw86340n8a85ge {}
2019-12-19T10:17:35+08:00 boudqumbcw08a43ri4aw86340n8a85ge Start Invoke
2019-12-19T10:17:35+08:00 boudrskojdn2sidrihc59obk9dchnf46 Report Invoke boudrskojdn2sidrihc59obk9dchnf46,Duration :7.50ms  BilledDuration: 100ms  Memory Size: 128 MB
2019-12-19T10:17:35+08:00 boudrskojdn2sidrihc59obk9dchnf46 End Invoke
2019-12-19T10:17:35+08:00 boudrskojdn2sidrihc59obk9dchnf46 {}
2019-12-19T10:17:35+08:00 boudrskojdn2sidrihc59obk9dchnf46 Start Invoke

```