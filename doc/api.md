@[TOC](接口)

 
<br>
 
## 1. 登录接口

### 请求头
- [ ] 请求 URL: `http://xxxx:8080/v1/tourist/login`
- [ ] 请求方式: `POST`
- [ ] 表头: `Content-Type: application/json`


### 请求参数
	
参数名|必选|类型|说明
-|-|-|-
`email`|是|`string`|登录邮箱
`password`|是| `string`|登录密码

### 示例
- [ ] 输入示例:
	```json
	{
		"email":"123@qq.com",
		"password":"123456789"
	}
	```

- [ ] 返回示例:
	```json
	{
		"err": "登录失败时才有",
		"msg": "登录成功时才有",
		"data": null
	}
	```
