# pbox:密码坐标

一个密码坐标由以下几个字段(属性)构成:

| 属性名称 | 类型                 | 说明                 |
| -------- | -------------------- | -------------------- |
| email    | string(EmailAddress) | 用户的 管理 账号     |
| domain1  | string(DomainName)   | 云公钥的域名         |
| domain2  | string(DomainName)   | 被托管账号所属的域名 |
| username | string               | 用户的 被托管账号    |
| scene    | string               | 密码的使用场景       |
| revision | int                  | 密码的版本           |
