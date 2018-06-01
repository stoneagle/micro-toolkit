# 工具集
## 基本功能
1. 提供各个平台的辅助功能

## 环境搭建
### 操作流程
1. 安装docker与docker-compose
2. 安装npm与ng
3. 安装golang与glide
4. 进入backend目录，执行glide install，安装backend依赖包
5. 进入frontend目录，执行npm install，安装frontend依赖包
6. 如果需要使用本地mysql，执行make init-db
7. 执行make build-golang构建backend开发镜像
8. 执行make run-web搭建服务端开发环境
9. 执行make run-ng搭建前端开发环境 

## 相关工具
### hack
1. 使用dockerfile打造基础和发布镜像
2. 使用docker-compose搭建开发环境
3. 使用glide完成第三方包管理
### backend
1. 框架gin
2. mysql数据库xorm
3. redis缓存go-redis
4. log日志zap
### frontend
1. angular5
2. vmware/clarity
