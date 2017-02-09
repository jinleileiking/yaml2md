# 设置录制视频定时任务(CreateRecordTask)

## 描述

设置录制视频定时任务

## 请求参数

| 名称 | 类型 | 是否必须 | 参数条件 | 默认值  | 描述 |
| --- | --- | --- | --- | --- | --- |
| Action | String | Y | CreateRecordTask | | |
| Version | String | Y | 2017-01-01 | | |
| UniqueName | String  | Y | [a-zA-Z0-9-] |  | 域名标识 | 
| App | String  | Y | [a-zA-Z0-9-] |  | app | 
| Pubdomain | String  | Y |  |  |  | 
| Stream | String  | Y |  |  |  | 
| StartUnixTime | String  | Y |  |  |  | 
| EndUnixTime | String  | Y |  |  |  | 
| Mp4VodEnable | String  | N |  |  |  | 
| Ks3FileNameM3U8 | String  | N |  |  |  | 
| Ks3FullPathMP4 | String  | N |  |  |  | 


## 返回参数

| 名称 | 类型 | 是否必须 |  描述 |
| --- | --- | --- |  --- |
| RecID | String  | Y | 录像id | 



## 请求示例

```
```

## 返回示例

```
```

