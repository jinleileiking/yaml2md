# 获取流生成ks3路径的文件名(GetRecordKs3Info)

## 描述

获取流生成ks3路径的文件名

## 请求参数

| 名称 | 类型 | 是否必须 | 参数条件 | 默认值  | 描述 |
| --- | --- | --- | --- | --- | --- |
| Action | String | Y | GetRecordKs3Info | | |
| Version | String | Y | 2017-01-01 | | |
| UniqueName | String  | Y | [a-zA-Z0-9-] |  | 域名标识 | 
| App | String  | Y | [a-zA-Z0-9-] |  | app | 
| Pubdomain | String  | Y |  |  |  | 
| Stream | String  | Y |  |  |  | 
| UnixTime | String  | Y |  |  |  | 
| Vdoid | String  | Y |  |  |  | 


## 返回参数

| 名称 | 类型 | 是否必须 |  描述 |
| --- | --- | --- |  --- |
| LiveFileNameM3U8 | String  | Y | 直播实时更新的m3u8文件名，参加命名方式，必须带vdoid | 
| PhraseFileNameMp4 | String  | Y |  | 



## 请求示例

```
```

## 返回示例

```
```

