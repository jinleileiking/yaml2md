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

{"brief"=>"获取域名级别配置",
 "method"=>"get",
 "version"=>#<Date: 2017-01-01 ((2457755j,0s,0n),+0s,2299161j)>,
 "input"=>
  [{"name"=>"UniqueName",
    "require"=>"[a-zA-Z0-9-]",
    "comment"=>"域名标识",
    "valid"=>true},
   {"name"=>"App", "require"=>"[a-zA-Z0-9-]", "comment"=>"app", "valid"=>true},
   {"name"=>"Pubdomain", "valid"=>true}],
 "output"=>
  [{"name"=>"PubdomainConfigs", "type"=>"Array", "comment"=>"推流域名的配置, 数组内容见下"},
   {"name"=>"App", "require"=>"[a-zA-Z0-9-]", "comment"=>"app"},
   {"name"=>"Pubdomain"},
   {"name"=>"Mp4VodEnable"},
   {"name"=>"HlsVodEnable"},
   {"name"=>"HlsBucket"},
   {"name"=>"Mp4Bucket"},
   {"name"=>"MergeEnable"},
   {"name"=>"LiveFileNameM3U8", "comment"=>"直播实时更新的m3u8文件名，参加命名方式，必须带vdoid"}],
 "action"=>"ListRecordConfigs",
 "detail"=>"获取域名级别配置"}
