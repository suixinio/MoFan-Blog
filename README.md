# mofan-blog

Gin + vue 一个个人博客。

## 已将实现功能

-[x] 简单的用户管理权限设置
-[x] 用户密码加密存储
-[x] 文章分类自定义
-[x] 列表分页
-[x] 图片上传七牛云
-[x] JWT 认证
-[x] 自定义日志功能
-[x] 跨域 cors 设置

## TODO LIST

-[ ] 增加Markdown编辑器，去掉HTML编辑器
-[ ] Github登陆进行文章评论功能
-[ ] 图片上传到自建的图床
-[ ] 一篇文章对应多个分类
-[ ] 书籍式的文章编辑
-[ ] 服务端渲染文章
-[ ] UI优化~~~
-[ ] UI优化~~~

UI优化是个细活~

## 运行

1. 安装依赖

```
go mod tidy
```

2. 初始化项目配置config.ini

```ini
./config/config.ini

[server]
AppMode = debug # debug 开发模式，release 生产模式
HttpPort = :3000 # 项目端口
JwtKey = 89js82js72 #JWT密钥，随机字符串即可

[database]
Db = mysql #这个参数暂时没有用到，代码写死是使用mysql
DbHost = 127.0.0.1 # 数据库地址
DbPort = 3306 # 数据库端口
DbUser = mofanbloog # 数据库用户名
DbPassWord = mofanbloog # 数据库用户密码
DbName = mofanbloog # 数据库名

[qiniu]
# 七牛储存信息
AccessKey = # AK
SecretKey = # SK
Bucket =
QiniuSever =
```

3. 使用gorm自动导入数据库


4. 启动项目

```shell
 go run main.go
```

此时，项目启动，你可以访问页面

```shell
首页
http://localhost:3000
后台管理页面
http://localhost:3000/admin

```

- 上线前请确认 `web` 文件夹下的axios请求地址

```js
// 在 web/admin/src/plugin/http.js 和 web/front/src/plugin/http.js 两个文件夹中,将 baseURL地址改为部署的服务器线上地址

axios.defaults.baseURL = 'http://localhost:3000/api/v1'

// 修改为

axios.defaults.baseURL = 'http://线上服务器ip或域名:3000/api/v1'
```

## 热重启

安装 gowatch ，启动之后，有文件更新会自动build

```
gowatch
```
