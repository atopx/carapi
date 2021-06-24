# ginhelper

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/yanmengfei/ginhelper)
![Github report](https://img.shields.io/badge/go%20report-A%2B-green)
[![GitHub stars](https://img.shields.io/github/stars/yanmengfei/ginhelper)](https://github.com/yanmengfei/ginhelper/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/yanmengfei/ginhelper)](https://github.com/yanmengfei/ginhelper/network)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/yanmengfei/ginhelper)
[![GitHub issues](https://img.shields.io/github/issues/yanmengfei/ginhelper)](https://github.com/yanmengfei/ginhelper/issues)


> 自动生成gin脚手架，包含orm、db连接池、docker运行、构建脚本等

## 工具使用

### 参数说明
```
NAME:
   ginhelper - Create a scaffold for the gin framework

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --name 项目名称, -n 项目名称    指定项目名称
   --output 项目路径, -o 项目路径  指定项目路径
   --remote GIT地址          指定GIT地址
   --docker                enable docker (default: false)
   --compose               use docker compose (default: false)
   --help, -h              show help (default: false)
   --version, -v           print the version (default: false)
```

### 创建工程
```bash
go get -u github.com/yanmengfei/ginhelper
ginhelper --name project --output /home/yanmengfei/opensource --docker --compose
```

## 结构说明
```
├── docker-compose.local.yaml // 本地docker启动 docker-compose -f docker-compose.local.yaml up
├── docker-compose.yaml // 镜像启动 docker-compose up
└── services // 服务目录
    └── app   // 代码
        ├── local.dockerfile // 本地docker启动调用
        ├── build.dockerfile // 构建docker镜像
        ├── go.mod
        ├── go.sum
        ├── main.go // 项目入口
        ├── config.yaml // 配置文件
        ├── api // 接口定义
        │   └── test.go 
        ├── config // 配置结构体定义
        │   └── config.go
        ├── core // 业务代码(CURD)
        ├── docs // swagger文档，运行`swag init`自动生成
        │   ├── docs.go
        │   ├── swagger.json
        │   └── swagger.yaml
        ├── libs // 工具库和外部依赖
        │   └── orm_logger.go
        ├── middleware // 中间件
        │   ├── cors.go
        │   ├── logger.go
        │   └── swagger.go
        ├── models // 模型定义
        │   └── user.go
        ├── public // 公共项
        │   ├── const.go // 公共常量
        │   └── variable.go // 公共变量
        ├── routers // 路由Group
        │   └── test.go
        ├── schemas // 参数验证
        │   ├── common.go
        │   └── response.go
        └── server // 服务初始化和配置
            ├── config.go // 加载配置文件
            ├── database.go
            ├── engine.go
            ├── logger.go
            └── sentry.go
```

