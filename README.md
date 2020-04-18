## DDD project-demo
### 项目结构
    ├── README.md
    ├── application                                       --- 应用层
    │   └── service
    │       └── leave.go
    ├── domain                                            --- 领域层
    │   ├── leave                                         --- 聚合
    │   │   ├── entity                                    --- 实体
    │   │   │   ├── approval.go
    │   │   │   ├── leave.go
    │   │   │   └── valueobject                           --- 值对象
    │   │   │       ├── ApprovalType.go
    │   │   │       ├── applicant.go
    │   │   │       └── approver.go
    │   │   ├── event                                     --- 领域事件
    │   │   │   ├── event_type.go
    │   │   │   └── leave.go
    │   │   ├── repository                                --- 仓储接口
    │   │   │   ├── facade
    │   │   │   │   └── leave.go
    │   │   │   ├── persistence                           --- CRUD
    │   │   │   │   └── mysql
    │   │   │   │       └── leave.go
    │   │   │   └── po                                    --- 持久化对象
    │   │   │       └── leave.go
    │   │   └── service
    │   │       ├── event_factory.go                      --- 工厂函数
    │   │       ├── leave.go
    │   │       └── leave_factory.go
    │   └── person                                        
    │       ├── entity
    │       │   └── person.go
    │       ├── repository
    │       │   ├── facade
    │       │   │   └── person.go
    │       │   ├── persistence
    │       │   │   └── mysql
    │       │   │       └── person.go
    │       │   └── po
    │       │       └── person.go
    │       └── service
    │           ├── person.go
    │           └── person_factory.go
    ├── infrastructure                                     --- 基础设施层
    │   ├── client                                         --- 调用外部接口
    │   ├── common                                         --- 公共代码
    │   │   ├── event
    │   │   │   └── base_event.go
    │   │   ├── po
    │   │   │   └── event.go
    │   │   └── response_code                              --- 返回码
    │   │       ├── code
    │   │       │   ├── grpc.pb.go
    │   │       │   └── leave.pb.go
    │   │       └── proto
    │   │           ├── grpc.proto
    │   │           └── leave.proto
    │   ├── config                                         --- 配置信息
    │   │   └── env.go
    │   └── util                                           --- 工具类
    │       ├── driver
    │       │   ├── kafka.go
    │       │   └── mysql.go
    │       ├── migrate
    │       │   └── migrate.go
    │       └── tools
    ├── interface                                          --- 用户接口层
    │   ├── dto                                            --- DTO对象定义
    │   │   ├── rest
    │   │   │   └── createLeave.go
    │   │   └── rpc
    │   │       └── leave.pb.go
    │   ├── factory                                        --- 工厂函数
    │   │   └── leave.go
    │   └── facade                                      --- 对外适配层
    │       ├── rest
    │       │   ├── handler
    │       │   │   ├── handler.go
    │       │   │   ├── leave.go
    │       │   │   └── ping.go
    │       │   └── router
    │       │       ├── middleware
    │       │       │   └── cors.go
    │       │       └── router.go
    │       └── rpc
    │           ├── client
    │           │   └── client.go
    │           ├── handler
    │           │   └── leave.go
    │           ├── proto
    │           │   └── leave.proto
    │           └── server
    │               └── server.go
    ├── main.go
    ├── docker-compose.yml
    ├── dockerfile
    ├── go.mod
    ├── go.sum
