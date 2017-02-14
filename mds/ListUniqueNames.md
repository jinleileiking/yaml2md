# 获取域名标识(ListUniqueNames)

## 描述

获取域名标识

## 请求参数

| 名称 | 类型 | 是否必须 | 参数条件 | 默认值  | 描述 |
| --- | --- | --- | --- | --- | --- |
| Action | String | Y | ListUniqueNames | | |
| Version | String | Y | 2017-01-01 | | |


## 返回参数

| 名称 | 类型 | 是否必须 |  描述 |
| --- | --- | --- |  --- |
| UniqueNames | String Array  | Y | 返回UniqueName数组 | 



## 请求示例

```
```

## 返回示例

```
```

{"brief"=>"获取域名标识",
 "method"=>"get",
 "version"=>#<Date: 2017-01-01 ((2457755j,0s,0n),+0s,2299161j)>,
 "input"=>nil,
 "output"=>
  [{"name"=>"UniqueNames",
    "type"=>"String Array",
    "comment"=>"返回UniqueName数组"}],
 "action"=>"ListUniqueNames",
 "detail"=>"获取域名标识"}
