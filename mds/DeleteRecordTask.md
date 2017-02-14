# 删除录制视频定时任务(DeleteRecordTask)

## 描述

删除录制视频定时任务

## 请求参数

| 名称 | 类型 | 是否必须 | 参数条件 | 默认值  | 描述 |
| --- | --- | --- | --- | --- | --- |
| Action | String | Y | DeleteRecordTask | | |
| Version | String | Y | 2017-01-01 | | |
| UniqueName | String  | Y | [a-zA-Z0-9-] |  | 域名标识 | 
| App | String  | Y | [a-zA-Z0-9-] |  | app | 
| Pubdomain | String  | Y |  |  |  | 
| Stream | String  | Y |  |  |  | 
| RecID | String  | Y | [a-zA-Z0-9-] |  | 录像id | 


## 返回参数

| 名称 | 类型 | 是否必须 |  描述 |
| --- | --- | --- |  --- |



## 请求示例

```
```

## 返回示例

```
```

{"brief"=>"删除录制视频定时任务",
 "method"=>"put",
 "version"=>#<Date: 2017-01-01 ((2457755j,0s,0n),+0s,2299161j)>,
 "input"=>
  [{"name"=>"UniqueName",
    "require"=>"[a-zA-Z0-9-]",
    "comment"=>"域名标识",
    "valid"=>true},
   {"name"=>"App", "require"=>"[a-zA-Z0-9-]", "comment"=>"app", "valid"=>true},
   {"name"=>"Pubdomain", "valid"=>true},
   {"name"=>"Stream", "valid"=>true},
   {"name"=>"RecID",
    "require"=>"[a-zA-Z0-9-]",
    "comment"=>"录像id",
    "valid"=>true}],
 "output"=>nil,
 "action"=>"DeleteRecordTask",
 "detail"=>"删除录制视频定时任务"}
