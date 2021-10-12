# learn-go
go build 编译二进制文件
go install 编译库文件 产生 x.a文件 //
go学习的日志
go mod init        初始化当前文件夹, 创建go.mod文件
go mod tidy        增加缺少的module，删除无用的module
go mod vendor      将依赖复制到vendor下
外网不通
        1.自己内网,没有依赖第三方库
        2.依赖内部其他团队提供的库 gitlab/gitee/
        2.依赖外网的第三方库
                a.编译环境 =>目标站点
                            =>acl
                            =>proxy
                            =>第三方包下载 上传到自己的库服务器

        gopath+vendor
   // overlapping:
    s = "888888888888888888"
    old = "888"
    new = "666"
    fmt.Println("overlapping: ")

    // n < 0 ,用 new 替换所有匹配上的 old；n=-1:  666666666666666666
    fmt.Println("n=-1: ", strings.Replace(s, old, new, -1 )) 

    // 不替换任何匹配的 old；n=0:  888888888888888888
    fmt.Println("n=0: ", strings.Replace(s, old, new, 0 )) 
  
    // 用 new 替换第一个匹配的 old；n=1:  666888888888888888
    fmt.Println("n=1: ", strings.Replace(s, old, new, 1 )) 

     // 用 new 替换前 5 个匹配的 old（实际多于 5 个）；n=5:  666666666666666888
    fmt.Println("n=5: ", strings.Replace(s, old, new, 5 ))

     // 用 new 替换前 7 个匹配上的 old（实际没那么多）；n=7:  666666666666666666
    fmt.Println("n=7: ", strings.Replace(s, old, new, 7 ))

    require "liwenzhou.com/q1mi/p2" v0.0.0   // v 0.0.0 指定版本
    replace "liwenzhou.com/q1mi/p2" => "../p2"

    go mod edit -replace=github.com/zzflygo/testmath@latest=github.com/imsilence/testmath@latest  //@latest 最后的版本 