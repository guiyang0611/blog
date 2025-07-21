

## Blog-个人博客系统

## 项目结构
blog/
- auth/ # 认证模块
  - auth #  认证模块
  - claim # 生成 JWT Token
- config/   # 配置文件           
     - config.dev.yaml
     - config.go  
- log/ # 日志模块 
     - logger.go
- models/ # 数据库模型结构体             
     - user.go # 用户模型
     - post.go # 文章模型
     - comment.go # 评论模型
- routers/ # 路由文件           
  - user.go
  - posts.go
  - comment.go
  - router.go 
- utils/ # 工具类文件
  - error.go # 统一异常处理
  - response.go # 统一返回结果
- go.mod # 项目依赖管理文件
- main.go  # 项目入口

## 初始化项目-依赖管理

- 初始化项目依赖管理: go mod init blog 
- 下载gin框架: go get -u github.com/gin-gonic/gin 
- 下载gorm框架: go get -u gorm.io/gorm
- 下载mysql驱动: go get -u gorm.io/driver/mysql
- 下载SqlX驱动:  go get -u github.com/jmoiron/sqlx
- 下载endless: go get github.com/fvbock/endless
- 下载viper: go get github.com/spf13/viper
- 下载zap: go get go.uber.org/zap
- 下载lumberjack.v2: go get gopkg.in/natefinch/lumberjack.v2
- 下载errors: github.com/pkg/errors
- 下载jwt :  go get "github.com/dgrijalva/jwt-go"
- 下载jwt :  go get "golang.org/x/crypto/bcrypt"



