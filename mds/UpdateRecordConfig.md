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

{"brief"=>"更新录像配置",
 "method"=>"post",
 "version"=>#<Date: 2017-01-01 ((2457755j,0s,0n),+0s,2299161j)>,
 "input"=>
  [{"name"=>"UniqueName",
    "require"=>"[a-zA-Z0-9-]",
    "comment"=>"域名标识",
    "valid"=>true},
   {"name"=>"App", "require"=>"[a-zA-Z0-9-]", "comment"=>"app", "valid"=>true},
   {"name"=>"Pubdomain", "valid"=>true},
   {"name"=>"Mp4VodEnable", "must"=>"N", "comment"=>"是否使能mp4, 这些参数至少带一个"},
   {"name"=>"MergeEnable", "must"=>"N"},
   {"name"=>"HlsBucket", "must"=>"N"},
   {"name"=>"Mp4Bucket", "must"=>"N"},
   {"name"=>"LiveFileNameM3U8",
    "must"=>"N",
    "comment"=>"直播实时更新的m3u8文件名，参加命名方式，必须带vdoid"}],
 "output"=>nil,
 "action"=>"UpdateRecordConfig",
 "detail"=>"更新录像配置"}
