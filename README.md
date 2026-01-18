# 存放我的go练习的代码库

代码很乱

go一般的目录
```go
project-name/
├── cmd/                        # 启动入口（可多个）
│   └── server/
│       └── main.go             # 程序入口
│
├── internal/                   # 仅限项目内部使用（Go 官方推荐）
│   ├── api/                    # HTTP 层（Controller / Handler）
│   │   ├── admin/
│   │   │   └── admin.go
│   │   ├── user/
│   │   │   └── user.go
│   │   └── health.go
│   │
│   ├── middleware/             # 中间件（横切逻辑）
│   │   ├── auth.go             # JWT / Session
│   │   ├── rbac.go             # 权限控制
│   │   ├── logger.go           # 日志
│   │   ├── recovery.go         # panic 恢复
│   │   └── ratelimit.go
│   │
│   ├── service/                # 业务逻辑层
│   │   ├── auth_service.go
│   │   ├── user_service.go
│   │   └── admin_service.go
│   │
│   ├── repository/             # 数据访问层（DAO）
│   │   ├── user_repo.go
│   │   └── role_repo.go
│   │
│   ├── model/                  # 数据模型
│   │   ├── user.go
│   │   ├── role.go
│   │   └── base.go
│   │
│   ├── router/                 # 路由聚合
│   │   └── router.go
│   │
│   ├── config/                 # 配置结构体
│   │   └── config.go
│   │
│   ├── bootstrap/              # 启动初始化
│   │   ├── app.go
│   │   ├── db.go
│   │   └── redis.go
│   │
│   └── pkgctx/                 # Context key 定义
│       └── context.go
│
├── pkg/                        # 可复用组件（可对外）
│   ├── jwt/                    # JWT 工具
│   │   ├── token.go
│   │   └── claims.go
│   ├── hash/                   # 密码哈希
│   ├── response/               # 统一返回结构
│   └── errors/                 # 业务错误码
│
├── configs/                    # 配置文件
│   ├── config.yaml
│   └── config.dev.yaml
│
├── scripts/                    # 脚本
│   └── migrate.sh
│
├── go.mod
└── go.sum

```