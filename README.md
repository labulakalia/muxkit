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