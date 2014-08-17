mod
===

HTTP Helper modules.

mod 处于开发期, 正式项目慎用.

风格
====

签名接口:
```go
/**
Interface for module
签名接口, 表示采用 mod 的接口风格.
*/
type Interface interface {
    GitHubTypepressMod()
}
```

所有的 module 都包含两个方法:

    签名 GitHubTypepressMod()
    执行 ServeHTTP

ServeHTTP 参数风格:
```go
/**
参数:
    最后两个参数类型是 `http.ResponseWriter`, `*http.Request`.
    之前的参数无限制.
返回:
    true  表示请求已完成, 应该停止响应
    false 表示请求未完成, 可以继续响应.
*/
ServeHTTP(v1 Type, vN TypeN, http.ResponseWriter, *http.Request) bool
```


LICENSE
=======
Copyright (c) 2014 The TypePress Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.