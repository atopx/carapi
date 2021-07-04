# carapi

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/yanmengfei/carapi)
![Github report](https://img.shields.io/badge/go%20report-A%2B-green)
[![GitHub stars](https://img.shields.io/github/stars/yanmengfei/carapi)](https://github.com/yanmengfei/carapi/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/yanmengfei/carapi)](https://github.com/yanmengfei/carapi/network)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/yanmengfei/carapi)
[![GitHub issues](https://img.shields.io/github/issues/yanmengfei/carapi)](https://github.com/yanmengfei/carapi/issues)


> 自动生成Go RestfulAPI企业级项目脚手架代码，基于最佳实践
> 
> 真正实现专注于业务逻辑的快速开发
> 
> - 没有框架 == 爬着开发
> - 有了框架 == 走着开发
> - 复用项目 == 跑了起来
> - carapi   == 开车狂飙

## TODO

### 框架

- [x] gin
- [x] fiber
- [ ] iris
- [ ] beego

### 数据库

- [x] posgresql
- [ ] mysql

## 工具使用

### 参数说明

```
NAME:
   carapi - Create a scaffold for the gin framework

USAGE:
   main.exe [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --name 项目名称, -n 项目名称    指定项目名称
   --frame 框架              选择框架 [ gin fiber iris beego ], 默认gin (default: "gin")
   --db 数据库                选择数据库 [ pgsql mysql ], 默认pgsql (default: "pgsql")
   --output 项目路径, -o 项目路径  指定项目路径
   --remote GIT地址          指定GIT地址
   --docker                enable docker (default: false)
   --compose               use docker compose (default: false)
   --help, -h              show help (default: false)
   --version, -v           print the version (default: false)
```

### 创建工程

```bash
go get -u github.com/yanmengfei/carapi
carapi --name newapp --frame fiber --output /Users/yanmengfe/opensource --docker --compose
```

## 结构说明

```shell
+--- .git   # git初始化
+--- .gitignore
+--- config.toml # 配置文件
+--- docker-compose.local.yaml # 本地docker启动
+--- docker-compose.yaml # compose部署
+--- release.sh  # 自动发布构建compose镜像脚本
+--- services 
|   +--- app # code根目录
|   |   +--- api # 接口
|   |   |   +--- task.go
|   |   |   +--- test.go
|   |   +--- build.dockerfile # 构建发布docker image
|   |   +--- config # 配置文件
|   |   |   +--- config.go
|   |   +--- docs # swagger
|   |   |   +--- docs.go
|   |   |   +--- swagger.json
|   |   |   +--- swagger.yaml
|   |   +--- go.mod
|   |   +--- go.sum
|   |   +--- local.dockerfile # 构建本地docker image
|   |   +--- main.go # 程序入口
|   |   +--- middleware  # 中间件
|   |   |   +--- cors.go
|   |   |   +--- logger.go
|   |   |   +--- swagger.go
|   |   +--- model  # 模型定义
|   |   |   +--- base.go
|   |   |   +--- task.go
|   |   +--- public # 公共资源
|   |   |   +--- const.go
|   |   |   +--- handle.go
|   |   |   +--- utils.go
|   |   +--- schema # 参数校验
|   |   |   +--- common.go
|   |   |   +--- task.go
|   |   +--- service # 逻辑代码
|   |   |   +--- task.go
|   |   +--- setup # 项目初始化配置
|   |   |   +--- config.go
|   |   |   +--- database.go
|   |   |   +--- engine.go
|   |   |   +--- logger.go
|   |   |   +--- router.go
```
