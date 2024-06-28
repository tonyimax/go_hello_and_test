package main //包名

import (
	"fmt"
	quoteV3 "rsc.io/quote/v3" //模块别名
	quoteV4 "rsc.io/quote/v4" //模块别名
	"testing"                 //导入测试模块
)

// 测试函数 规则：Test + 包中要测试的函数名，如Hello
// 参数，测试类testing.T指针
func TestGetHello(t *testing.T) {
	tStr := "Hello,world."
	if r := getHello(); r != tStr {
		t.Errorf("测试Hello()函数失败: r=%q,tStr=%q", r, tStr)
	}
}

func TestGetEmpty(t *testing.T) {
	fmt.Println("===>", quoteV3.HelloV3())
	fmt.Println("===>", quoteV4.Opt())
	empty := getEmpty()
	if empty == "" {
		t.Errorf("测试HelloEmpty()函数失败,传入字符串为空")
	}
}
