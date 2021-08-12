# golang-uuid-loader
golang uuid-shellcode加载器,分离执行,可直接把shellcode写入程序。

# 相关项目
golang-uuid：https://github.com/Ne0nd0g/go-shellcode/blob/master/cmd/UuidFromString/main.go  （uuid）

### 如何编译(加密shellcode)
```
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -ldflags "-s -w -H=windowsgui" -o code.exe encode.go
```
### 如何编译(加载器)
```
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -ldflags "-s -w -H=windowsgui" -o aescodeloade.exe main.go
```
### 使用方法
code.exe  用于生成加密的shellcode，使用方法：用cs生成一个64位的raw格式的payload，然后用code.exe加密生成一个加密code和一个key

aescode.exe 加载器 -s 跟加密的shellcode -k 跟解密的key


用cs生成一个64位的raw格式的payload
2021/08/12
