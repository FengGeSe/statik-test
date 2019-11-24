# Statik-test

根据[statik](https://github.com/rakyll/statik)项目，制作的deamo.

将static目录下的index.html文件，打包进入代码。运行http文件服务器，提供网页服务。

ps: 部署方便， 一个二进制文件包含了html等静态资源。

### Usage

##### 生成静态文件

运行如下命令 将static目录所有文件生成statik/statik.go文件。

```go generate```

##### 运行

```go run main.go```

访问 http://127.0.0.1

