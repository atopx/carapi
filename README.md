# ginhelper

```
├── docker-compose.local.yaml    // 本地docker启动 docker-compose -f docker-compose.local.yaml up
├── docker-compose.yaml  // 镜像启动 docker-compose up
└── services      // 服务目录
    └── app       // 代码
        ├── local.dockerfile  // 本地docker启动调用
        ├── build.dockerfile  // 构建docker镜像
        ├── go.mod
        ├── go.sum
        ├── main.go // 项目入口
        ├── config.yaml  // 配置文件
        ├── api    // 接口定义
        │   └── test.go 
        ├── config  // 配置结构体定义
        │   └── config.go
        ├── core    // 业务代码
        ├── docs    // swagger文档，运行`swag init`自动生成
        │   ├── docs.go
        │   ├── swagger.json
        │   └── swagger.yaml
        ├── libs  // 工具库和外部依赖
        │   └── org_logger.go
        ├── middleware // 中间件
        │   ├── cors.go
        │   ├── logger.go
        │   └── swagger.go
        ├── models // 模型定义
        │   └── user.go
        ├── public // 公共项
        │   ├── const.go
        │   └── variable.go
        ├── routers // 路由组
        │   └── test.go
        ├── schemas // 参数验证
        │   ├── common.go
        │   └── response.go
        └── server // 服务初始化
            ├── config.go
            ├── database.go
            ├── engine.go
            ├── logger.go
            └── sentry.go
```

以下文件夹内，每个路由组新建一个文件
  - schemas
  - core
  - api
  - routers