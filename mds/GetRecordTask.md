# 获取某个任务情况(GetRecordTask)

## 描述

获取某个任务情况

## 请求参数

| 名称 | 类型 | 是否必须 | 参数条件 | 默认值  | 描述 |
| --- | --- | --- | --- | --- | --- |
| Action | String | Y | GetRecordTask | | |
| Version | String | Y | 2017-01-01 | | |
| RecID | String  | Y | [a-zA-Z0-9-] |  | 录像id | 


## 返回参数

| 名称 | 类型 | 是否必须 |  描述 |
| --- | --- | --- |  --- |
| UniqueName | String  | Y | 域名标识 | 
| App | String  | Y | app | 
| Pubdomain | String  | Y |  | 
| Stream | String  | Y |  | 
| RecStatus | String  | Y | 录像状态 1--计划将要录制任务，2--实时正在录制的任务，3--已经完成的录制任务<成功/失败>，0--全部 | 
| RecType | String  | Y | 录像类型 0:全部 1:定时 2:小视频 | 
| StartUnixTime | String  | Y |  | 
| EndUnixTime | String  | Y |  | 
| Ks3FullPathMp4 | String  | Y | ks3存储全路径，包括目录和文件名，命名方式:{App}{StreamName}{UnixTime}{YY}{MM}{DD}{HH}{mm}{ss}{vdoid} | 
| Ks3FileNameM3u8 | String  | Y |  | 



## 请求示例

```
```

## 返回示例

```
```

