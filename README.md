# muxkit


- example
  - cmd    // 程序启动目录
    - example
      - main.go
  - middle // 中间件
    - middle.go
  - config // 配置 mysql redis other   
    - config.go 
  - deploy // 部署相关
    - TODO
  模块
  - example1
    - proto // proto文件 用来生成用户库
      - example1.proto
      - example1.pb.go
      - example1.http.pb.go
    - db    // 数据库
      - db.go // 使用默认的配置 并且可以覆盖
    - server // 业务逻辑
      - server.go
    - dao    // 数据库操作



## 路由
url里的路由比url参数的值优先级高
/test/{id}?id=10


type HelloRequest struct {
    HelloID   uint64 `json:"hello_id"`
}

请求URL /test/12?id=30
最终的获取到的参数Hello.HelloID值为12，30会被覆盖，因为url里的参数值比？后面的优先级高
