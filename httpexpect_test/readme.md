### 运行命令
* 在分支master运行测试

```sh 
make do_test master
```

* 在分支master运行集成测试

```sh 
make do_integration_test master
```

* 在分支master运行单元测试

```sh 
make do_unit_test master
```

* 停止当前测试

```sh 
make stop_test 
```

* 停止当前集成测试

```sh 
make stop_integration_test
```

* 停止当前单元测试

```sh 
make stop_unit_test 
```

### 备注
1. 单元测试的开始为启动脚本或所选分支有代码提交的时候，检查提交的间隔为执行完测试后五分钟
2. 集成测试的开始为启动脚本或所选分支有代码更新或上次测试通过，检查提交的间隔为执行完测试后1分钟
3. 停止测试的意义为不会进行新一轮的测试，不会在测试执行中即```go test```运行中强制停止
4. 使用的是 [httpexpect](https://github.com/gavv/httpexpect) 库，Concise, declarative, and easy to use end-to-end HTTP and REST API testing for Go (golang).

### 建议
1. 单元测试以功能命名文件夹名，以测试功能命名文件，一个文件一个test方法，多个case使用```t.run()```。如测试 user 模块：
```
|--unit_test
   |--user
      |--get_user_test.go
      |--get_user_list_test.go
      |--post_user_test.go
```
2. 集成测试同上
3. 测试的目的是提前发现低级bug。具体集成测试与单元测试自行Google。
