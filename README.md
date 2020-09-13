# 概述
OaPms 是GO语言实现的多应用版本的Oauth2权限系统。该项目参考了[zeus-admin](https://github.com/bullteam/zeus-admin)和[gin-admin](https://github.com/LyricTian/gin-admin)的实现。使用bee工具的beegopro功能自动生成前后端，加以简单改造完成该项目。

# 介绍
OaPms使用了GO，React技术，Gin，Ant Deign Pro4框架。使得我们的系统更加易于部署，开发和维护。
功能点
* [x] Oauth2服务，兼容gitlab oauth2
* [x] Casbin RBAC权限
* [x] 全局级别Google身份验证器
* [ ] Oauth网关
* [ ] PMS SDK
* [ ] 钉钉、企业微信组织结构
* [ ] 数据权限

## 运行
```
go get github.com/beego/bee
安装后端
make install
运行后端
make go
运行前端
make ant

```
