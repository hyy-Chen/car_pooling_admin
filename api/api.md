# api 文档介绍

## 错误返回

### 登录（token过期）错误返回

- **状态码**： 401

- **响应体：** application/json

  ```json
  {
    "status": "error",
    "message": "令牌错误"
  }
  ```

### token更新返回

- **状态码**： 401

- **响应体：** application/json

  ```json
  {
    "status": "error",
    "message": "Token updates",
    "data": {
     	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjMsInVzZXJuYW1lIjoiam9obiJ9.7K8LsGdZg5e0WmFJzT8T2h9sNpM1eI6L7cXWnD5gW3E"  
     }
  }
  ```


### 参数错误返回

- **状态码**： 400

- **响应体：** application/json

  ```json
  {
    "status": "error",
    "message": "参数错误"
  }
  ```

### 服务器内部错误返回

- **状态码**： 400

- **响应体：** application/json

  ```json
  {
    "status": "error",
    "message": "服务器错误"
  }
  ```



## 登录模块

### 登录接口

该接口用于用户进行身份验证获取令牌。

接口信息

- **路径：** `/api/login`
- **请求方法：** POST
- **内容类型：** application/json
- **身份验证：** 不需要身份验证

请求参数

| 参数名称 | 类型   | 必需 | 描述   |
| -------- | ------ | ---- | ------ |
| account  | string | 是   | 用户名 |
| password | string | 是   | 密码   |

响应

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

## 个人信息管理模块
### 获取个人信息接口

该接口用于获取当前用户个人信息

**接口信息**

- **路径**：`api/admin/information`

- **方法**：GET

- **身份验证：** 请求头：`Authorization:  <token>` <token>信息替换成浏览器缓存里的令牌

**响应**


- **状态码：** 200
- **响应体：** application/json

```json
{
  "status": "success",
  "message": "信息获取成功",
  "data": {
    "admin": {
          "id": 123,
          "schoolId": "your_school_id",
          "name": "John Doe",
          "nickName": "JD",
          "account": "johndoe",
          "createPerson": 456,
          "role": 1,
          "createTime": "2022-01-01T12:34:56Z",
          "updateTime": "2022-01-02T10:20:30Z"
    }
  }
}
```

### 更新用户信息

该接口用于更新用户信息模块

**接口信息**

- **路径**：`api/admin/information`
- **方法**：PUT 

- **内容类型：** application/json

- **身份验证：** 请求头：`Authorization:  <token>` <token>信息替换成浏览器缓存里的令牌

- **请求参数** 

| 参数名称     | 类型   | 必需 | 描述     |
| ------------ | ------ | ---- | -------- |
| id           | int64  | 是   | 用户ID   |
| name         | string | 否   | 姓名     |
| nickName     | string | 否   | 昵称     |
| password     | string | 否   | 密码     |

需要更新哪一部分信息就传过来对应参数(建议一次性更新一个信息)

**响应**

- **状态码：** 200
- **响应体：** application/json

```json
{
  "status": "success",
  "message": "信息更新成功"
}
```

## 管理员信息管理模块（超管可用:role=1）

### 获取管理员信息列表

该接口获取部分管理员信息,客户端需要传递起始下标值和长度值作为参数，后端将返回相应范围内的管理员信息列表。

**接口信息**

- **路径**：`api/admin/informations`
- **方法**：GET 

- **内容类型：** application/json

- **身份验证：** 请求头：`Authorization:  <token>` <token>信息替换成浏览器缓存里的令牌

- **请求参数** 

| 参数名 | 类型 | 是否必填 | 描述       |
| ------ | ---- | -------- | ---------- |
| start  | 整数 | 是       | 起始下标值 |
| length | 整数 | 是       | 长度值     |

**响应**

- **状态码：** 200
- **响应体：** application/json

```json
{
  "status": "success",
  "message": "信息获取成功",
  "data": {
    	"admin_list":[
            {
                  "id": 123,
                  "schoolId": "your_school_id",
                  "name": "John Doe",
                  "nickName": "JD",
                  "account": "johndoe",
                  "createPerson": 456,
                  "role": 1,
                  "createTime": "2022-01-01T12:34:56Z",
                  "updateTime": "2022-01-02T10:20:30Z"
            }
        ],
      	"total": len(admin_list)
	}
}
```

 

### 重置指定管理员密码

该接口通过对应管理员id将管理员密码重置。

**接口信息**

- **路径**：`api/admin/informations/password`
- **方法**：PUT

- **内容类型：** application/json

- **身份验证：** 请求头：`Authorization:  <token>` <token>信息替换成浏览器缓存里的令牌

- **请求参数** 

| 参数名称    | 类型   | 必需 | 描述       |
| ----------- | ------ | ---- | ---------- |
| id          | int64  | 是   | 用户ID     |
| newPassword | string | 是   | 更新后密码 |

**响应**

- **状态码：** 200
- **响应体：** application/json

```json
{
  "status": "success",
  "message": "密码重置成功"
}
```

### 删除指定管理员信息

通过指定管理员id删除管理员的操作，客户端需要传递用户ID作为参数，后端将根据ID删除相应的用户信息并返回成功或失败的响应。

**接口信息**

- **路径**：`api/admin/informations`
- **方法**：DELETE

- **内容类型：** application/json

- **身份验证：** 请求头：`Authorization:  <token>` <token>信息替换成浏览器缓存里的令牌

- **请求参数** 

| 参数名称    | 类型   | 必需 | 描述       |
| ----------- | ------ | ---- | ---------- |
| id          | int64  | 是   | 用户ID     |

**响应**

- **状态码：** 200
- **响应体：** application/json

```json
{
  "status": "success",
  "message": "删除成功"
}
```

### 根据ID搜索管理员信息

该接口用于通过ID的模糊匹配查询管理员信息。客户端需要传递管理员ID的部分匹配值作为参数，后端将根据匹配值查询相应的管理员信息并返回管理员信息列表。

**接口信息**

- **路径**：`api/admin/informations`
- **方法**：DELETE

- **内容类型：** application/json

- **身份验证：** 请求头：`Authorization:  <token>` <token>信息替换成浏览器缓存里的令牌

- **请求参数** 

| 参数名称 | 类型  | 必需 | 描述   |
| -------- | ----- | ---- | ------ |
| id       | int64 | 是   | 用户ID |
| start  | 整数 | 是       | 起始下标值 |
| length | 整数 | 是       | 长度值     |

**响应**

- **状态码：** 200
- **响应体：** application/json

```json
{
  "status": "success",
  "message": "搜索成功",
  "data": {
      "admin_list":[
          {
              "id": 123,
              "schoolId": "your_school_id",
              "name": "John Doe",
              "nickName": "JD",
              "account": "johndoe",
              "createPerson": 456,
              "role": 1,
              "createTime": "2022-01-01T12:34:56Z",
              "updateTime": "2022-01-02T10:20:30Z"
        }
      ],
      "total": len(admin_list)
  }
}
```
