# 获取域名级别配置(ListRecordConfigs)

## 描述

获取域名级别配置

## 请求参数

| 名称 | 类型 | 是否必须 | 参数条件 | 默认值  | 描述 |
| --- | --- | --- | --- | --- | --- |
| Action | String | Y | ListRecordConfigs | | |
| Version | String | Y | 2017-01-01 | | |
| UniqueName | String  | Y | [a-zA-Z0-9-] |  | 域名标识 | 
| App | String  | Y | [a-zA-Z0-9-] |  | app | 
| Pubdomain | String  | Y |  |  |  | 


## 返回参数

| 名称 | 类型 | 是否必须 |  描述 |
| --- | --- | --- |  --- |
| PubdomainConfigs | Array  | Y | 推流域名的配置, 数组内容见下 | 
| App | String  | Y | app | 
| Pubdomain | String  | Y |  | 
| Mp4VodEnable | String  | Y |  | 
| HlsVodEnable | String  | Y |  | 
| HlsBucket | String  | Y |  | 
| Mp4Bucket | String  | Y |  | 
| MergeEnable | String  | Y |  | 
| LiveFileNameM3U8 | String  | Y | 直播实时更新的m3u8文件名，参加命名方式，必须带vdoid | 



## 请求示例

```
```

## 返回示例

```
```

