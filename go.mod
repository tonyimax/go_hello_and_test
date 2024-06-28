//终端输入 go mod init com.test/main 生成 go.mod
module com.test/main //生成的模块名

go 1.22.2 //生成的go版本号

//终端输入 go get rsc.io/quote 生成的远程模块引用
require (
	golang.org/x/text v0.15.0 // indirect
	rsc.io/quote v1.5.2
	rsc.io/sampler v1.3.1 // indirect
)

require rsc.io/quote/v3 v3.1.0 //quote V3版本
require rsc.io/quote/v4 v4.0.1 //quote V4版本
