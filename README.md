# :package:UserAgent :bookmark:v1.0 :tada::tada::tada:

✨ ___练手项目___

Golang 开发的随机获取 UserAgent 的库和 API

## Installation

``` sh
$ cd mygoapp
$ go get github.com/msadg/UserAgent
```

## Use

``` go
package main
// ...
import "github.com/msadg/UserAgent
// ...
```

## :memo: Data

> ./data/user-agent.json

## :wrench: UserAgent Library

* 随机获取一条 UserAgent

    ``` go
    UserAgent.Rand() string
    ```

* 获取所有的 UserAgent

    ``` go
    UserAgent.All() []string
    ```

* 所有浏览器名

    ``` go
    UserAgent.ListBrowsers() []string
    ```

* 按设备浏览器随机获取一条 UserAgent

    ``` go
    // chrome
    // opera
    // firefox
    // internet explorer
    // safari
    // microsoft edge
    // iPhone
    // android
    // other
    UserAgent.RandBs(name string) string
    ```

* 获取某个浏览器所属的所有 UserAgent

    ``` go
    UserAgent.BrowserAll(name string) []string
    ```

## :mag:UserAgent API

> api/server/user_agent_api.go

:rocket:Run

``` sh
go run ./api/server [<-p :9580>]
```

* 全部 UserAgent

    ``` url
    /api
    ```

* 同上

    ``` url
    /api?ua=all
    ```

* 随机一条

    ``` url
    /api?ua=rand
    ```

* 获取所有 Chrome浏览器 的 UserAgent

    ``` url
    /api?ua=chrome
    ```

* 同上

    ``` url
    /api?ua=chrome&dt=all
    ```

* 随机获取一条 Chrome 浏览器的 UserAgent

    ``` url
    /api?ua=chrome&dt=rand
    ```

* 获取5条Chrome浏览器的 UserAgent

    ``` url
    /api?ua=chrome&dt=5
    ```

:bulb:Url 访问链接参数

``` text
/api[?ua={ua}[&dt={dt}]]

{ua}: rand | {browserName} | all

    {browserName}: 
        chrome
        opera
        firefox
        ie
        safari
        edge
        iPhone
        android
        other

{dt}: rand/{count}/all

    {count}: uint
```
