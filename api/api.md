# api 文档介绍

## 登录接口

该接口用于用户进行身份验证和登录。

### 接口信息

- **路径：** `/admin/login`
- **请求方法：** POST
- **内容类型：** application/json
- **身份验证：** 不需要身份验证

### 请求参数

| 参数名称 | 类型   | 必需 | 描述   |
| -------- | ------ | ---- | ------ |
| account  | string | 是   | 用户名 |
| password | string | 是   | 密码   |

### 响应

#### 成功响应

- **状态码：** 200

- **响应体：** application/json

  ```json
  {
    "status": "success",
    "message": "登录成功",
    "data": {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjMsInVzZXJuYW1lIjoiam9obiJ9.7K8LsGdZg5e0WmFJzT8T2h9sNpM1eI6L7cXWnD5gW3E"
    }
  }
  ```

#### 错误响应

- **状态码：** 401
- **响应体：** application/json

```json
{
  "status": "error",
  "message": "用户名或密码错误"
}
```
