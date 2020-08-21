
## 简介
构建中。

## 技术栈
- [ ] `react`
- [ ] `go`
	- `gin`
	- `websocket`
	- `jwt`
	- `gorm`
	- `imaging`
- [ ] `mysql`
- [ ] `redis`
- [ ] `grafana`
- [ ] `nginx`


## 初步实现的功能 (未测试)
- [ ] 邮箱注册
- [ ] 登录
- [ ] 修改密码
- [ ] 上传、获取头像
- [ ] 上传、获取图片
- [ ] 关注
- [ ] 消息发送
- [ ] 实时对话


## 开发指南
- [ ] `react`
	- [`api`文档](https://github.com/Lxy417165709/curve/tree/master/doc/api.md)
- [ ] `grafana`
	- [进入grafana主页](https://github.com/Lxy417165709/curve/tree/master/grafana/readme.md)

## 风格
- [ ] 交互: `restful`(正在学习)
- [ ] 命名:
	- 标识符命名: 大驼峰、小驼峰
	- 文件命名: 蛇形命名
	- 标签命名: 蛇形命名
	- 包命名: 小写字母 



## 经验收获
- [ ] `go`
	- 出现错误时，不仅仅要进行处理，更要输出错误，这样在出现 `bug` 时，就能很快找出原因。
	- 输出日志时，一定要记录输出代码的位置。

## 常用指令
- [ ] `go`
	- `go mod init <模块名>`: 初始化模块。
	- `go mod tidy`: 同步模块。
- [ ] `redis`
	- `redis-server`: 开启 `redis` 服务器。
	- `redis-cli -h 127.0.0.1 -p 6379`: 连接 `redis` 服务器。
	- `keys *`: 查看 `redis` 服务器的所有键。

## 资料
- [ ] `go`
	- [gorm出现incorrect datetime value '0000-0-0 00:00:00' for column问题](https://www.jianshu.com/p/3a2a7c61cce1)
	- [gorm文档](http://gorm.book.jasperxu.com/)
- [ ] `redis`
	- [在Mac上安装redis](https://www.cnblogs.com/DI-DIAO/p/12588078.html)
- [ ] 规范
	- [接口文档编写示例](https://www.jianshu.com/p/a5ceb0611125)