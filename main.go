package main //包名
import (
	"fmt"
	"rsc.io/quote"            //quote默认版本
	quoteV3 "rsc.io/quote/v3" //quote指定v3版本，并重新命名模块别名为quoteV3
)

// Hello 方法
func getHello() string {
	fmt.Println(quote.Hello()) //调用quote模块Hello方法
	quoteV3.Concurrency()      //使用quoteV3模块别名调用模块方法Concurrency
	return "Hello,world."      //返回字符串
}

// getEmpty 方法
func getEmpty() string {
	fmt.Println(quote.Go()) //调用quote模块Hello方法
	return ""               //返回字符串
}

func main() {
	getHello()
	getEmpty()
}
