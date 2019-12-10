## GOPATH的理解
GOPATH其实算是GO的一个败笔，之前开发的项目都需要放置到GOPATH的src目录下面，这个
开发起来就优点坑了啊，然后，GO的目的是想不同的项目使用的GOPATH不同，但是这个可以
根据版本来控制啊，怎么使用目录了
## Go的module的使用
module是go的一个包依赖管理工具，类似于maven，看来go还是朝向版本管理工具屈服了
不过这个要配置点东西
```
    GO111MODULE=on
    GOPROXY=https://goproxy.io 
```
## GO的module的使用
```
    go mod init xxxx 
    :初始化一个项目，xxxx是项目名字，记住，xxxx会被写入到go.mod文件的兽行
    所以，这个如果你不给名字，那就是这种形式:github.com/github的用户名/xxxx(项目名字)
    
    之后直接go run main.go就行了，他会直接下载东西
```

## 字典类型
```cassandraql
说文解字 · 龙龛手鉴 · 康熙字典 · 干禄字书 · 五经文字 · 九经字样 · 方言 · 字汇 · 玉篇 · 正字通 · 广雅 · 尔雅 · 小尔雅 · 埤雅 · 释名 · 一切经音义
```