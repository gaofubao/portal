# portal
Portal

## 常用命令
```shell
# 启动服务
go run portal.go -f etc/portal-api.yaml

# 生成 api 服务 
goctl api go --api api/portal.api --dir . --style=go_zero

# 生成 api 文档
goctl api doc --dir api/portal.api --o api/portal.md
```

## 参考文档
- [官方文档](https://go-zero.dev/cn/docs/introduction)
