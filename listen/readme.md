用于监控 18080 端口，通过码云的 webHook 的 post 请求达到效果。
restart-goVps.log 文件会记录执行过程

流程：
1、监听18080端口
2、请求到达18080端口，执行restart.sh
3、编译信息会在文件restart-goVps.log上。

运行：
go build -o listen