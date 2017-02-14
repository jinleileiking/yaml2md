# 获取流级别的录像任务情况(ListRecordTasks)

## 描述

获取流级别的录像任务情况

## 请求参数

| 名称 | 类型 | 是否必须 | 参数条件 | 默认值  | 描述 |
| --- | --- | --- | --- | --- | --- |
| Action | String | Y | ListRecordTasks | | |
| Version | String | Y | 2017-01-01 | | |
| UniqueName | String  | Y | [a-zA-Z0-9-] |  | 域名标识 | 
| App | String  | Y | [a-zA-Z0-9-] |  | app | 
| Pubdomain | String  | Y |  |  |  | 
| Stream | String  | Y |  |  |  | 
| RecType | String  | N |  | 0 全部 | 录像类型 0:全部 1:定时 2:小视频 | 
| OrderTime | Int  | N |  | 0 从最新的开始 | 排序类型 0:从前向后 1:从后向前 | 
| Marker | String  | N |  |  | 起始位置 | 
| Limit | String  | N |  |  | 限制数量 | 
| StartUnixTime | String  | N |  |  |  | 
| EndUnixTime | String  | N |  |  |  | 


## 返回参数

| 名称 | 类型 | 是否必须 |  描述 |
| --- | --- | --- |  --- |
| UniqueName | String  | Y | 域名标识 | 
| App | String  | Y | app | 
| Pubdomain | String  | Y |  | 
| Stream | String  | Y |  | 
| OrderTime | Int  | Y | 排序类型 0:从前向后 1:从后向前 | 
| RecType | String  | Y | 录像类型 0:全部 1:定时 2:小视频 | 
| Total | Int  | Y | 总数 | 
| Count | Int  | Y | 数量 | 
| Recs | Array  | Y | 返回任务数组，具体见下 | 
| RecID | String  | Y | 录像id | 
| RecStatus | String  | Y | 录像状态 1--计划将要录制任务，2--实时正在录制的任务，3--已经完成的录制任务<成功/失败>，0--全部 | 
| RecType | String  | Y | 这个任务的录制类型 | 
| StartUnixTime | String  | Y |  | 
| EndUnixTime | String  | Y |  | 
| Ks3FullPathMp4 | String  | Y | ks3存储全路径，包括目录和文件名，命名方式:{App}{StreamName}{UnixTime}{YY}{MM}{DD}{HH}{mm}{ss}{vdoid} | 
| Ks3FullPathM3u8 | String  | Y | ks3存储全路径，包括目录和文件名，命名方式:{StreamName}{UnixTime}{YY}{MM}{DD}{HH}{mm}{ss} | 



## 请求示例

```
```

## 返回示例

```
```

{"brief"=>"获取流级别的录像任务情况",
 "version"=>#<Date: 2017-01-01 ((2457755j,0s,0n),+0s,2299161j)>,
 "input"=>
  [{"name"=>"UniqueName",
    "require"=>"[a-zA-Z0-9-]",
    "comment"=>"域名标识",
    "valid"=>true},
   {"name"=>"App", "require"=>"[a-zA-Z0-9-]", "comment"=>"app", "valid"=>true},
   {"name"=>"Pubdomain", "valid"=>true},
   {"name"=>"Stream", "valid"=>true},
   {"name"=>"RecType",
    "must"=>"N",
    "default"=>"0 全部",
    "comment"=>"录像类型 0:全部 1:定时 2:小视频"},
   {"name"=>"OrderTime",
    "must"=>"N",
    "default"=>"0 从最新的开始",
    "type"=>"Int",
    "comment"=>"排序类型 0:从前向后 1:从后向前"},
   {"name"=>"Marker", "must"=>"N", "comment"=>"起始位置"},
   {"name"=>"Limit", "must"=>"N", "comment"=>"限制数量"},
   {"name"=>"StartUnixTime", "must"=>"N"},
   {"name"=>"EndUnixTime", "must"=>"N"}],
 "output"=>
  [{"name"=>"UniqueName", "require"=>"[a-zA-Z0-9-]", "comment"=>"域名标识"},
   {"name"=>"App", "require"=>"[a-zA-Z0-9-]", "comment"=>"app"},
   {"name"=>"Pubdomain"},
   {"name"=>"Stream"},
   {"name"=>"OrderTime", "type"=>"Int", "comment"=>"排序类型 0:从前向后 1:从后向前"},
   {"name"=>"RecType", "comment"=>"录像类型 0:全部 1:定时 2:小视频"},
   {"name"=>"Total", "type"=>"Int", "comment"=>"总数"},
   {"name"=>"Count", "type"=>"Int", "comment"=>"数量"},
   {"name"=>"Recs", "type"=>"Array", "comment"=>"返回任务数组，具体见下"},
   {"name"=>"RecID", "require"=>"[a-zA-Z0-9-]", "comment"=>"录像id"},
   {"name"=>"RecStatus",
    "comment"=>"录像状态 1--计划将要录制任务，2--实时正在录制的任务，3--已经完成的录制任务<成功/失败>，0--全部"},
   {"name"=>"RecType", "comment"=>"这个任务的录制类型"},
   {"name"=>"StartUnixTime"},
   {"name"=>"EndUnixTime"},
   {"name"=>"Ks3FullPathMp4",
    "comment"=>
     "ks3存储全路径，包括目录和文件名，命名方式:{App}{StreamName}{UnixTime}{YY}{MM}{DD}{HH}{mm}{ss}{vdoid}"},
   {"name"=>"Ks3FullPathM3u8",
    "comment"=>
     "ks3存储全路径，包括目录和文件名，命名方式:{StreamName}{UnixTime}{YY}{MM}{DD}{HH}{mm}{ss}"}],
 "action"=>"ListRecordTasks",
 "method"=>"GET",
 "detail"=>"获取流级别的录像任务情况"}
