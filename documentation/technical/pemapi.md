# PEM-API

PEM-API 是 PasswordBox 前端-后端 之间的调用协议.
这个协议基于 PEM 文件格式, 并参考了 HTTP-based RESTful API 的设计风格.

## 接口列表

以下是目前支持的接口:

| Method | Path                                         | 功能                    | 备注 |
| ------ | -------------------------------------------- | ----------------------- | ---- |
| GET    | /banks/root/                                 | 获取  ROOT-Bank 信息    |      |
| POST   | /banks/root/init                             | 初始化 ROOT-Bank        |      |
| POST   | /banks/root/objects/                         | 创建新对象 @Root-Bank   |      |
| GET    | /banks/root/objects/{object_id}              | 获取对象信息 @Root-Bank |      |
| GET    | /banks/root/refs/                            | 获取 refs 列表          |      |
|        |                                              |                         |      |
| GET    | /banks/{user_email_addr}/                    | 获取 用户-Bank 信息     |      |
| POST   | /banks/{user_email_addr}/init                | 初始化 用户-Bank        |      |
| POST   | /banks/{user_email_addr}/objects/            | 创建新对象 @User-Bank   |      |
| GET    | /banks/{user_email_addr}/objects/{object_id} | 获取对象信息 @User-Bank |      |
| GET    | /banks/{user_email_addr}/refs/               | 获取 refs 列表          |      |


