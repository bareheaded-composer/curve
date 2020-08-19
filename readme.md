
## 简介
构建中。

## 技术栈
- [ ] `go`
	- `gin`
	- `websocket`
	- `jwt`
	- `gorm`
	- `imaging`
- [ ] `mysql`
- [ ] `redis`

## 风格
- [ ] 交互: `restful`(正在学习)
- [ ] 命名:
	- 标识符命名: 大驼峰、小驼峰
	- 文件命名: 蛇形命名
	- 标签命名: 蛇形命名
	- 包命名: 小写字母 

## 常用指令
- [ ] `go`
	- `go mod init <模块名>`: 初始化模块。
	- `go mod tidy`: 同步模块。
- [ ] `redis`
	- `redis-server`: 开启 `redis` 服务器。
	- `redis-cli -h 127.0.0.1 -p 6379`: 连接 `redis` 服务器。
	- `keys *`: 查看 `redis` 服务器的所有键。

## 交互

### `go`后端响应结构
```go
type Response struct {
	Err  string      `json:"err"`  // 错误消息，如格式错误
	Msg  string      `json:"msg"`  // 一般消息，如登录成功
	Data interface{} `json:"data"` // 响应数据
}
```
### 接口
待写


## 初步实现的功能 (未测试)
- [ ] 邮箱注册
- [ ] 登录
- [ ] 修改密码
- [ ] 上传、获取头像
- [ ] 上传、获取图片
- [ ] 关注
- [ ] 消息发送
- [ ] 实时对话


## 资料
- [ ] `go`
	- [gorm出现incorrect datetime value '0000-0-0 00:00:00' for column问题](https://www.jianshu.com/p/3a2a7c61cce1)
	- [gorm文档](http://gorm.book.jasperxu.com/)
- [ ] `redis`
	- [在Mac上安装redis](https://www.cnblogs.com/DI-DIAO/p/12588078.html)