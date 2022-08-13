# :package:UserAgent :bookmark:v1.0 :tada::tada::tada:

✨ ___练手项目___

Golang 开发的随机获取 UserAgent 的库和 API

## :memo: 数据文件

> ./data/user-agent.json

## :wrench: UserAgent Library

* 随机获取一条 UserAgent

    ``` go
    userAgent.Rand() string
    ```

* 获取所有的 User-Agent

    ``` go
    userAgent.All() []string
    ```

* 所有浏览器名

    ``` go
    userAgent.ListBrowsers() []string
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
    userAgent.RandBs(name string) string
    ```

* 获取某个浏览器所属的所有 UserAgent

    ``` go
    userAgent.BrowserAll(name string) []string
    ```

## :mag:UserAgent API

> api/server/user_agent_api.go

:rocket:启动服务

``` sh
go run ./api/server -p :9580
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

``` http
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

{dt}: rand/{count}/all

    {count}: uint
```
