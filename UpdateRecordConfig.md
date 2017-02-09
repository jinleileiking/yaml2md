# 更新录像配置(UpdateRecordConfig)

## 描述

更新录像配置

## 请求参数

| 名称 | 类型 | 是否必须 | 参数条件 | 默认值  | 描述 |
| --- | --- | --- | --- | --- | --- |
| Action | String | Y | UpdateRecordConfig | | |
| Version | String | Y | 2017-01-01 | | |
| UniqueName | String  | Y | [a-zA-Z0-9-] |  | 域名标识 | 
| App | String  | Y | [a-zA-Z0-9-] |  | app | 
| Pubdomain | String  | Y |  |  |  | 
| Mp4VodEnable | String  | N |  |  | 是否使能mp4, 这些参数至少带一个 | 
| MergeEnable | String  | N |  |  |  | 
| HlsBucket | String  | N |  |  |  | 
| Mp4Bucket | String  | N |  |  |  | 
| LiveFileNameM3U8 | String  | N |  |  | 直播实时更新的m3u8文件名，参加命名方式，必须带vdoid | 


## 返回参数

| 名称 | 类型 | 是否必须 |  描述 |
| --- | --- | --- |  --- |



## 请求示例

```
```

## 返回示例

```
```

