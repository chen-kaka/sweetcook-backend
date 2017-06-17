
# 高性能通用web服务器(基于gin)

### gin

github:

[gin-gonic/gin](https://github.com/gin-gonic/gin)

DOC:

[gin-GoDoc](https://godoc.org/github.com/gin-gonic/gin)

依赖：

参考glide.yaml

使用glide进行管理:

        curl https://glide.sh/get | sh

        or download at: https://glide.sh/

安装glide依赖：

        glide install

install的时候可能会出错，需要配置代理：

        sudo vi /etc/profile
        export http_proxy=http://127.0.0.1:55945
        export https_proxy=$http_proxy
        export ftp_proxy=$http_proxy
        export rsync_proxy=$http_proxy
        export no_proxy="localhost,127.0.0.1,localaddress,.localdomain.com"

        source /etc/profile

热加载：

    go get github.com/codegangsta/gin
    gin -h

启动：

开发环境：

    sh run-dev.sh

测试环境：

    sh run-test.sh

测试：

在test目录下新建单元测试用例，然后cd到该目录，执行go test即可。

使用gowatch进行热编译：

    go get github.com/silenceper/gowatch