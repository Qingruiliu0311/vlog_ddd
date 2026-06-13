# vlog_ddd

基于领域驱动设计（DDD）的博客内容管理系统，使用 Go + Gin + GORM 构建。

## 技术栈

| 组件 | 技术 |
|------|------|
| 语言 | Go 1.25+ |
| Web 框架 | Gin v1.12 |
| ORM | GORM v1.31 + MySQL |
| 认证 | Token（UUID，数据库存储） |
| 密码加密 | bcrypt |
| 参数校验 | go-playground/validator |

## 项目结构

```
vlog_ddd/
├── main.go              # 入口，路由注册与依赖初始化
├── blog/                # 博客领域
│   ├── model.go         # 领域模型
│   ├── interface.go     # 服务接口
│   ├── enum.go          # 枚举（草稿/已发布）
│   ├── api/api.go       # HTTP Handler
│   ├── impl/impl.go     # 业务逻辑实现
│   └── docs/table.sql   # 建表语句
├── user/                # 用户领域
│   ├── model.go
│   ├── interface.go
│   ├── impl/impl.go
│   └── docs/table.sql
├── token/               # Token 认证领域
│   ├── model.go
│   ├── interface.go
│   ├── api/api.go
│   ├── impl/impl.go
│   └── doc/table.sql
├── middleware/          # 认证中间件
└── exception/           # 统一异常处理
```

## 快速开始

### 环境依赖

- Go 1.21+
- MySQL 8.0+

### 1. 创建数据库

```sql
CREATE DATABASE vlog CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
```

### 2. 初始化表结构

```bash
mysql -u root -p vlog < blog/docs/table.sql
mysql -u root -p vlog < user/docs/table.sql
mysql -u root -p vlog < token/doc/table.sql
```

> 默认数据库连接：`root:123456@localhost:3306/vlog`，可在 `main.go` 中修改。

### 3. 安装依赖并启动

```bash
go mod download
go run main.go
```

服务启动在 `http://localhost:8080`。

## API 文档

所有需要认证的接口须在 Header 中携带：`Authorization: Bearer <access_token>`

### 认证

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/v1/token` | 登录，颁发 Token |
| POST | `/api/v1/token/validate` | 校验 Token 有效性 |

**登录示例：**

```json
POST /api/v1/token
{
  "username": "user@example.com",
  "password": "yourpassword"
}
```

**响应：**

```json
{
  "access_token": "...",
  "refresh_token": "...",
  "access_token_expired_at": "...",
  "refresh_token_expired_at": "..."
}
```

- Access Token 有效期：**1 天**
- Refresh Token 有效期：**7 天**

### 博客

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|
| POST | `/api/v1/blog` | 创建博客 | 是 |
| GET  | `/api/v1/blog` | 查询博客列表 | 是 |

**创建博客示例：**

```json
POST /api/v1/blog
{
  "title": "标题",
  "content": "正文内容",
  "summary": "摘要",
  "tag": ["Go", "DDD"],
  "catelog": "技术"
}
```

**查询博客支持的过滤参数：**

| 参数 | 说明 |
|------|------|
| `stage` | 状态（0=草稿，1=已发布） |
| `keyword` | 标题/内容关键词 |
| `catelog` | 分类 |
| `tag` | 标签 |

## 架构设计

项目遵循 DDD 分层架构，每个领域（blog、user、token）内部结构一致：

```
interface.go   →  定义服务契约（接口）
model.go       →  领域模型与请求/响应结构
impl/impl.go   →  业务逻辑实现（依赖注入 DB）
api/api.go     →  HTTP 层，调用 Service 接口
```

异常处理统一由 `exception/` 包管理，包含标准化错误码与 HTTP 状态码映射。

## License

MIT
